package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	grant "github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"go.uber.org/zap"
)

type groupBuilder struct {
	resourceType *v2.ResourceType
	client       *msgraphsdkgo.GraphServiceClient
}

func newGroupResource(ctx context.Context, group models.Groupable) (*v2.Resource, error) {
	displayName := group.GetDisplayName()
	if displayName == nil {
		return nil, wrapError(nil, "group does not have a display name")
	}
	groupName := group.GetMailNickname()
	if groupName == nil {
		return nil, wrapError(nil, "group does not have a mail nickname")
	}
	groupId := group.GetId()
	if groupId == nil {
		return nil, wrapError(nil, "group does not have an id")
	}

	profile := map[string]interface{}{
		"name":     *displayName,
		"group_id": *groupId,
	}

	groupTraits := []rs.GroupTraitOption{
		rs.WithGroupProfile(profile),
	}

	resource, err := rs.NewGroupResource(*displayName, groupResourceType, *groupId, groupTraits)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (g *groupBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return groupResourceType
}

func (g *groupBuilder) List(ctx context.Context, _ *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	collection, err := g.client.Groups().Get(ctx, &groups.GroupsRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.GroupsRequestBuilderGetQueryParameters{
			Top: &resourcePageSize,
		},
	})
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get groups")
	}

	iterator, err := msgraphgocore.NewPageIterator[models.Groupable](collection, g.client.GetAdapter(), models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to create group page iterator")
	}

	var innerErr error
	var resources []*v2.Resource
	err = iterator.Iterate(ctx, func(group models.Groupable) bool {
		resource, err := newGroupResource(ctx, group)
		if err != nil {
			innerErr = wrapError(err, "failed to create group resource")

			return false
		}

		resources = append(resources, resource)

		return true
	})
	if innerErr != nil {
		return nil, "", nil, innerErr
	}
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate group collection")
	}

	return resources, "", nil, nil
}

func (g *groupBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	assigmentOptions := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("Member of %s group", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s group %s", resource.DisplayName, memberEntitlement)),
	}

	entitlement := ent.NewAssignmentEntitlement(resource, memberEntitlement, assigmentOptions...)
	rv = append(rv, entitlement)

	return rv, "", nil, nil
}

func (g *groupBuilder) Grants(ctx context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	collection, err := g.client.Groups().ByGroupId(resource.Id.Resource).Members().GraphUser().Get(ctx, &groups.ItemMembersGraphUserRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemMembersGraphUserRequestBuilderGetQueryParameters{
			Top: &resourcePageSize,
		},
	})
	if err != nil {
		return nil, "", nil, wrapError(err, "Failed to get users")
	}

	iterator, err := msgraphgocore.NewPageIterator[models.Userable](collection, g.client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to create user page iterator")
	}

	var innerErr error
	var grants []*v2.Grant
	err = iterator.Iterate(ctx, func(user models.Userable) bool {
		userResource, err := newUserResource(ctx, user)
		if err != nil {
			innerErr = wrapError(err, "failed to create user resource")

			return false
		}

		g := grant.NewGrant(resource, memberEntitlement, userResource.Id)

		grants = append(grants, g)

		return true
	})
	if innerErr != nil {
		return nil, "", nil, innerErr
	}
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate user collection")
	}

	return grants, "", nil, nil
}

func (g *groupBuilder) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	if principal.Id.ResourceType != userResourceType.Id {
		err := fmt.Errorf("baton-ms365: only users can be granted to groups")

		l.Warn(
			err.Error(),
			zap.String("principal_id", principal.Id.Resource),
			zap.String("principal_type", principal.Id.ResourceType),
		)
	}

	reference := models.NewReferenceCreate()
	oDataId := getODataId(principal.Id.Resource)
	reference.SetOdataId(&oDataId)
	err := g.client.Groups().ByGroupId(entitlement.Resource.Id.Resource).Members().Ref().Post(ctx, reference, nil)
	if err != nil {
		err := wrapError(err, "failed to grant user to group")

		l.Error(
			err.Error(),
			zap.String("group_id", entitlement.Resource.Id.Resource),
			zap.String("user_id", principal.Id.Resource),
		)
	}

	return nil, nil
}

func (g *groupBuilder) Revoke(ctx context.Context, grant *v2.Grant) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	entitlement := grant.Entitlement
	principal := grant.Principal

	if principal.Id.ResourceType != userResourceType.Id {
		err := fmt.Errorf("baton-ms365: only users can be revoked from groups")

		l.Warn(
			err.Error(),
			zap.String("principal_id", principal.Id.Resource),
			zap.String("principal_type", principal.Id.ResourceType),
		)
	}

	err := g.client.Groups().ByGroupId(entitlement.Resource.Id.Resource).Members().ByDirectoryObjectId(principal.Id.Resource).Ref().Delete(ctx, nil)
	if err != nil {
		err := wrapError(err, "failed to revoke user from group")

		l.Error(
			err.Error(),
			zap.String("group_id", entitlement.Resource.Id.Resource),
			zap.String("user_id", principal.Id.Resource),
		)
	}

	return nil, nil
}

func newGroupBuilder(client *msgraphsdkgo.GraphServiceClient) *groupBuilder {
	return &groupBuilder{
		resourceType: groupResourceType,
		client:       client,
	}
}
