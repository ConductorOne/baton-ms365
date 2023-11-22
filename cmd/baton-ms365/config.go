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

	ClientId string `mapstructure:"client-id"`
	TenantId string `mapstructure:"tenant-id"`

	// One of the following must be set
	ClientCertificatePath string `mapstructure:"client-certificate-path"`
	ClientSecret          string `mapstructure:"client-secret"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.ClientId == "" {
		return fmt.Errorf("client-id is required")
	}

	if cfg.TenantId == "" {
		return fmt.Errorf("tenant-id is required")
	}

	if cfg.ClientCertificatePath == "" && cfg.ClientSecret == "" {
		return fmt.Errorf("either client-certificate-path or client-secret is required")
	}

	return nil
}

func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("client-id", "", "Client ID")
	cmd.PersistentFlags().String("tenant-id", "", "Tenant ID")
	cmd.PersistentFlags().String("client-certificate-path", "", "Path to client certificate file")
	cmd.PersistentFlags().String("client-secret", "", "Client Secret")
}

func (c *config) GetMS365AuthenticationMethod() connector.AuthenticationMethod {
	if c.ClientCertificatePath != "" {
		return connector.WithClientCertificate(c.ClientCertificatePath)
	}

	return connector.WithClientSecret(c.ClientSecret)
}
