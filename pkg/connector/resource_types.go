package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
)

// The user resource type is for all user objects from the database.
var (
	userResourceType = &v2.ResourceType{
		Id:          "user",
		DisplayName: "User",
		Description: "MS365 user",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
		Annotations: annotations.New(
			&v2.SkipEntitlements{},
		),
	}
	roleResourceType = &v2.ResourceType{
		Id:          "role",
		DisplayName: "Role",
		Description: "Admin role of MS365",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
	}
	groupResourceType = &v2.ResourceType{
		Id:          "group",
		DisplayName: "Group",
		Description: "MS365 group",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_GROUP},
	}
	licenseResourceType = &v2.ResourceType{
		Id:          "license",
		DisplayName: "License",
		Description: "MS365 subscribed SKU (license)",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_LICENSE_PROFILE},
		Annotations: annotations.New(
			&v2.SkipGrants{},
			&v2.OptInRequired{},
			&v2.CapabilityPermissions{
				Permissions: []*v2.CapabilityPermission{
					{
						Permission: "LicenseAssignment.Read.All",
					},
				},
			},
		),
	}
)
