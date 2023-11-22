package connector

import (
	"context"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphsdkgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

type userBuilder struct {
	resourceType *v2.ResourceType
	client       *msgraphsdkgo.GraphServiceClient
}

func newUserResouce(ctx context.Context, user models.Userable) (*v2.Resource, error) {
	firstName := user.GetGivenName()
	if firstName == nil {
		firstName = user.GetDisplayName()
	}

	userName := *user.GetUserPrincipalName()

	profile := map[string]interface{}{
		"first_name":   user.GetGivenName(),
		"last_name":    user.GetSurname(),
		"display_name": user.GetDisplayName(),
		"email":        user.GetMail(),
		"login":        userName,
		"user_id":      user.GetId(),
	}

	var status v2.UserTrait_Status_Status
	if user.GetAccountEnabled() != nil && *user.GetAccountEnabled() {
		status = v2.UserTrait_Status_STATUS_ENABLED
	} else {
		status = v2.UserTrait_Status_STATUS_DISABLED

	}

	userTraits := []rs.UserTraitOption{
		rs.WithUserProfile(profile),
		rs.WithUserLogin(userName),
		rs.WithStatus(status),
	}

	resource, err := rs.NewUserResource(
		*user.GetDisplayName(),
		userResourceType,
		user.GetId(),
		userTraits,
	)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (o *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (o *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	collection, err := o.client.Users().Get(ctx, &users.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Top: &resourcePageSize,
		},
	})
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get users collection")
	}

	iterator, err := msgraphsdkgocore.NewPageIterator[models.Userable](collection, o.client.GetAdapter(), models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate users collection")
	}

	var innerErr error
	var resources []*v2.Resource
	err = iterator.Iterate(ctx, func(user models.Userable) bool {
		resource, err := newUserResouce(ctx, user)
		if err != nil {
			innerErr = wrapError(err, "failed to create user resource")

			return false
		}

		resources = append(resources, resource)

		return true
	})
	if innerErr != nil {
		return nil, "", nil, innerErr
	}
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to iterate users collection")
	}

	return nil, "", nil, nil
}

// Entitlements always returns an empty slice for users.
func (o *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *userBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newUserBuilder(client *msgraphsdkgo.GraphServiceClient) *userBuilder {
	return &userBuilder{
		resourceType: userResourceType,
		client:       client,
	}
}
