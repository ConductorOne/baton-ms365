package config

import (
	"fmt"

	"github.com/conductorone/baton-ms365/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	Ms365ClientId = field.StringField(
		"ms365-client-id",
		field.WithDescription("Client ID"),
		field.WithRequired(true),
		field.WithDisplayName("MS365 Client ID"),
	)
	Ms365TenantId = field.StringField(
		"ms365-tenant-id",
		field.WithDescription("Tenant ID"),
		field.WithRequired(true),
		field.WithDisplayName("MS365 Tenant ID"),
	)
	Ms365ClientSecret = field.StringField(
		"ms365-client-secret",
		field.WithDescription("Client Secret"),
		field.WithIsSecret(true),
		field.WithDisplayName("MS365 Client Secret"),
	)
	Ms365ClientCertificatePath = field.StringField(
		"ms365-client-certificate-path",
		field.WithDescription("Path to client certificate file"),
		field.WithDisplayName("MS365 Client Certificate Path"),
	)
	BaseURLField = field.StringField(
		"base-url",
		field.WithDescription("Override the Microsoft Graph API URL (for testing)"),
		field.WithHidden(true),
	)

	// FieldRelationships defines relationships between the fields listed in
	// Config that can be automatically validated.
	FieldRelationships = []field.SchemaFieldRelationship{}
)

//go:generate go run ./gen
var Config = field.NewConfiguration([]field.SchemaField{
	Ms365ClientId,
	Ms365TenantId,
	Ms365ClientSecret,
	Ms365ClientCertificatePath,
	BaseURLField,
})

// ValidateConfig is run after the configuration is loaded, and should return an
// error if it isn't valid.
func ValidateConfig(cfg *Ms365) error {
	if cfg.Ms365ClientId == "" {
		return fmt.Errorf("ms365-client-id is required")
	}

	if cfg.Ms365TenantId == "" {
		return fmt.Errorf("ms365-tenant-id is required")
	}

	if cfg.Ms365ClientCertificatePath == "" && cfg.Ms365ClientSecret == "" {
		return fmt.Errorf("either ms365-client-certificate-path or ms365-client-secret is required")
	}

	return nil
}

// GetAuthenticationMethod returns the appropriate authentication method based on config.
func GetAuthenticationMethod(cfg *Ms365) connector.AuthenticationMethod {
	if cfg.Ms365ClientCertificatePath != "" {
		return connector.WithClientCertificate(cfg.Ms365ClientCertificatePath)
	}

	return connector.WithClientSecret(cfg.Ms365ClientSecret)
}
