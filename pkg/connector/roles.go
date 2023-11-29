package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	"github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type roleBuilder struct {
	resourceType *v2.ResourceType
	client       *msgraphsdkgo.GraphServiceClient
}

func newRoleResource(ctx context.Context, role models.DirectoryRoleable) (*v2.Resource, error) {
	displayName := role.GetDisplayName()
	if displayName == nil {
		return nil, wrapError(nil, "role does not have a display name")
	}
	roleId := role.GetId()
	if roleId == nil {
		return nil, wrapError(nil, "role does not have an id")
	}

	profile := map[string]interface{}{
		"name":    *displayName,
		"role_id": *roleId,
	}

	roleTraits := []rs.RoleTraitOption{
		rs.WithRoleProfile(profile),
	}

	resource, err := rs.NewRoleResource(*displayName, roleResourceType, *roleId, roleTraits)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (r *roleBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return roleResourceType
}

func (r *roleBuilder) List(ctx context.Context, _ *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	collection, err := r.client.DirectoryRoles().Get(ctx, nil)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get roles")
	}

	iterator, err := msgraphgocore.NewPageIterator[models.DirectoryRoleable](collection, r.client.GetAdapter(), models.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to create role iterator")
	}

	var innerErr error
	var resources []*v2.Resource
	err = iterator.Iterate(ctx, func(role models.DirectoryRoleable) bool {
		resource, err := newRoleResource(ctx, role)
		if err != nil {
			innerErr = err
			return false
		}

		resources = append(resources, resource)
		return true
	})
	if innerErr != nil {
		return nil, "", nil, innerErr
	}
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate roles")
	}

	return resources, "", nil, nil
}

func (r *roleBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	assigmentOptions := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("Assigned to %s role", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s role %s", resource.DisplayName, assignedEntitlement)),
	}
	e := ent.NewAssignmentEntitlement(resource, assignedEntitlement, assigmentOptions...)
	rv = append(rv, e)

	return rv, "", nil, nil
}

func (r *roleBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	collection, err := r.client.DirectoryRoles().ByDirectoryRoleId(resource.Id.Resource).Members().Get(ctx, nil)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get role")
	}

	iterator, err := msgraphgocore.NewPageIterator[models.Userable](collection, r.client.GetAdapter(), models.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to create role members iterator")
	}

	var innerErr error
	var grants []*v2.Grant
	err = iterator.Iterate(ctx, func(user models.Userable) bool {
		userResource, err := newUserResource(ctx, user)
		if err != nil {
			innerErr = wrapError(err, "failed to create user resource")

			return false
		}

		g := grant.NewGrant(resource, assignedEntitlement, userResource.Id)

		grants = append(grants, g)

		return true
	})
	if innerErr != nil {
		return nil, "", nil, innerErr
	}
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate role members")
	}

	return grants, "", nil, nil
}

func newRoleBuilder(client *msgraphsdkgo.GraphServiceClient) *roleBuilder {
	return &roleBuilder{
		resourceType: roleResourceType,
		client:       client,
	}
}
