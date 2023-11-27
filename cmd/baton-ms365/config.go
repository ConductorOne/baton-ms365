package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-ms365/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options

	MS365ClientId string `mapstructure:"ms365-client-id"`
	MS365TenantId string `mapstructure:"ms365-tenant-id"`

	// One of the following must be set
	MS365ClientCertificatePath string `mapstructure:"ms365-client-certificate-path"`
	MS365ClientSecret          string `mapstructure:"ms365-client-secret"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.MS365ClientId == "" {
		return fmt.Errorf("ms365-client-id is required")
	}

	if cfg.MS365TenantId == "" {
		return fmt.Errorf("ms365-tenant-id is required")
	}

	if cfg.MS365ClientCertificatePath == "" && cfg.MS365ClientSecret == "" {
		return fmt.Errorf("either ms365-client-certificate-path or ms365-client-secret is required")
	}

	return nil
}

func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("ms365-client-id", "", "Client ID")
	cmd.PersistentFlags().String("ms365-tenant-id", "", "Tenant ID")
	cmd.PersistentFlags().String("ms365-client-secret", "", "Client Secret")
	cmd.PersistentFlags().String("ms365-client-certificate-path", "", "Path to client certificate file")
}

func (c *config) GetMS365AuthenticationMethod() connector.AuthenticationMethod {
	if c.MS365ClientCertificatePath != "" {
		return connector.WithClientCertificate(c.MS365ClientCertificatePath)
	}

	return connector.WithClientSecret(c.MS365ClientSecret)
}
