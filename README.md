![Baton Logo](./docs/images/baton-logo.png)

# `baton-ms365` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-ms365.svg)](https://pkg.go.dev/github.com/conductorone/baton-ms365) ![main ci](https://github.com/conductorone/baton-ms365/actions/workflows/main.yaml/badge.svg)

`baton-ms365` is a connector for MS365 built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It works with [Microsoft Graph API](https://learn.microsoft.com/en-gb/graph/api/overview?view=graph-rest-1.0).

Check out [Baton](https://github.com/conductorone/baton) to learn more about the project in general.

# Prerequisites

Connector use Client credentials OAuth 2.0 flow. As described [here](https://learn.microsoft.com/en-us/entra/identity-platform/v2-oauth2-client-creds-grant-flow) connector exchange client id and client secret (or certificate) for access token. To obtain these credentials, you have to register an app as described in [this tutorial](https://learn.microsoft.com/en-gb/graph/auth-register-app-v2?context=graph%2Fapi%2F1.0&view=graph-rest-1.0#add-credentials). You can decide if you want to use client secret or certificate.

After you have obtained client id and secret, you can use them with connector. You can do this by setting `BATON_MS365_CLIENT_ID` and `BATON_MS365_CLIENT_SECRET` (or `BATON_MS365_CLIENT_CERTIFICATE_PATH`) environment variables or by passing them as flags to baton-ms365 command. Also you have to set `BATON_MS365_TENANT_ID`.

# Getting Started

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-ms365

BATON_MS365_TENANT_ID=uuid BATON_MS365_CLIENT_ID=another-uuid BATON_MS365_CLIENT_SECRET=secret baton-ms365
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_MS365_TENANT_ID=uuid BATON_MS365_CLIENT_ID=another-uuid BATON_MS365_CLIENT_SECRET=secret ghcr.io/conductorone/baton-ms365:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-ms365/cmd/baton-ms365@main

BATON_MS365_TENANT_ID=uuid BATON_MS365_CLIENT_ID=another-uuid BATON_MS365_CLIENT_SECRET=secret baton-ms365
baton resources
```

# Data Model

`baton-ms365` will fetch information about the following MS365 resources:

- Users
- Groups
- Roles

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-ms365` Command Line Usage

```
baton-ms365

Usage:
  baton-ms365 [flags]
  baton-ms365 [command]

Available Commands:
  capabilities       Get connector capabilities
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --client-id string                       The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string                   The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string                            The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                                   help for baton-ms365
      --log-format string                      The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string                       The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
      --ms365-client-certificate-path string   Path to client certificate file
      --ms365-client-id string                 Client ID
      --ms365-client-secret string             Client Secret
      --ms365-tenant-id string                 Tenant ID
  -p, --provisioning                           This must be set in order for provisioning actions to be enabled. ($BATON_PROVISIONING)
  -v, --version                                version for baton-ms365

Use "baton-ms365 [command] --help" for more information about a command.
```