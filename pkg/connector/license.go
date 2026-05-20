package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type licenseBuilder struct {
	resourceType *v2.ResourceType
	client       *msgraphsdkgo.GraphServiceClient
}

func newLicenseResource(sku models.SubscribedSkuable) (*v2.Resource, error) {
	skuID := sku.GetSkuId()
	if skuID == nil {
		return nil, wrapError(nil, "subscribed sku does not have a skuId")
	}
	skuIDStr := skuID.String()

	displayName := skuIDStr
	if partNumber := sku.GetSkuPartNumber(); partNumber != nil && *partNumber != "" {
		displayName = *partNumber
	}

	var purchased, consumed int64
	if prepaid := sku.GetPrepaidUnits(); prepaid != nil {
		if enabled := prepaid.GetEnabled(); enabled != nil {
			purchased = int64(*enabled)
		}
	}
	if c := sku.GetConsumedUnits(); c != nil {
		consumed = int64(*c)
	}

	entitlementID := fmt.Sprintf("%s:%s:%s", licenseResourceType.Id, skuIDStr, licenseEntitlementAssigned)

	traitOpts := []rs.LicenseProfileTraitOption{
		rs.WithLicenseName(displayName),
		rs.WithLicenseSeats(purchased, consumed),
		rs.WithLicenseEntitlementIDs(entitlementID),
	}

	resource, err := rs.NewResource(
		displayName,
		licenseResourceType,
		skuIDStr,
		rs.WithLicenseProfileTrait(traitOpts...),
		rs.WithExternalID(&v2.ExternalId{Id: skuIDStr}),
	)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (l *licenseBuilder) ResourceType(_ context.Context) *v2.ResourceType {
	return licenseResourceType
}

func (l *licenseBuilder) List(ctx context.Context, _ *v2.ResourceId, _ *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	collection, err := l.client.SubscribedSkus().Get(ctx, nil)
	if err != nil {
		return nil, "", nil, wrapError(err, "failed to get subscribed skus")
	}

	var resources []*v2.Resource
	for _, sku := range collection.GetValue() {
		resource, err := newLicenseResource(sku)
		if err != nil {
			return nil, "", nil, wrapError(err, "failed to create license resource")
		}
		resources = append(resources, resource)
	}

	return resources, "", nil, nil
}

func (l *licenseBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	opts := []ent.EntitlementOption{
		ent.WithGrantableTo(userResourceType),
		ent.WithDescription(fmt.Sprintf("Assigned %s license", resource.DisplayName)),
		ent.WithDisplayName(fmt.Sprintf("%s license %s", resource.DisplayName, licenseEntitlementAssigned)),
	}
	e := ent.NewAssignmentEntitlement(resource, licenseEntitlementAssigned, opts...)
	return []*v2.Entitlement{e}, "", nil, nil
}

func (l *licenseBuilder) Grants(ctx context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newLicenseBuilder(client *msgraphsdkgo.GraphServiceClient) *licenseBuilder {
	return &licenseBuilder{
		resourceType: licenseResourceType,
		client:       client,
	}
}
