# `github.com/hashicorp/go-azure-sdk/microsoft-graph`

This SDK is an opinionated Go SDK for Microsoft Graph, primarily intended to be used in the [Terraform AzureAD Provider](https://github.com/hashicorp/terraform-provider-azuread).

More information on this, and our Resource Manager SDK, can be found in [the main `README.md` file](../README.md).

## Overview

This SDK is published as a Go module, and makes use of the `github.com/hashicorp/go-azure-sdk/sdk` module for authentication, cloud environment configuration, and helper types. Packages within this SDK are organised in the structure `./microsoft-graph/{serviceName}/{apiVersion}/{resourceName}`, where:

* `{serviceName}` is a distinct Microsoft Graph service such as `applications`, `groups`, `policies` etc.
* `{apiVersion}` corresponds to the v1.0 or beta Microsoft Graph API version, and is named either `stable` (for v1.0) or `beta`.
* `{resourceName}` is a Terraform-centric grouping of endpoints that represent the same logical resource, and which comprise one or more distinct operations.

Additionally, the majority of models and constants (i.e. enums) can be found in the `stable` and `beta` packages within the `common-types` directory. Some resources also define their own models specifically for request/response payloads that are unique to that resource.

## Getting Started

We have designed this SDK to be as easy-to-use as possible, and you can get started with minimal boilerplate.

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Minute)
	defer cancel()

	environment := environments.AzurePublic()
	credentials := auth.Credentials{
		Environment:                       *environment,
		EnableAuthenticatingUsingAzureCLI: true,
		// you can also authenticate using a Service Principal, Managed Identity, OIDC etc.
	}
	log.Printf("[DEBUG] Obtaining Credentials..")
	authorizer, err := auth.NewAuthorizerFromCredentials(ctx, credentials, environment.MicrosoftGraph)
	if err != nil {
		log.Fatalf("building authorizer from credentials: %+v", err)
	}

	log.Printf("[DEBUG] Building Application Client..")
	client, err := application.NewApplicationClientWithBaseURI(environment.MicrosoftGraph)
	if err != nil {
		log.Fatalf("building Application client: %+v", err)
	}
	client.Client.SetAuthorizer(authorizer)

	log.Printf("[DEBUG] Creating application..")
	payload := stable.Application{
		DisplayName: nullable.Value("My Application"),
		SignInAudience: nullable.Value("AzureADMultipleOrgs"),
	}

	options := application.DefaultCreateApplicationOperationOptions()

	resp, err := client.CreateApplication(ctx, payload, options)
	if err != nil {
		log.Fatalf("Creating application: %v", err)
	}

	if resp.Model == nil {
		log.Fatalf("Creating application: response was empty")
	}
	if resp.Model.Id == nil {
		log.Fatalf("Creating application: no object ID returned")
	}

	log.Printf("[DEBUG] Application created with object ID: %s", *resp.Model.Id)
}
```

For more usage examples, please refer to the [HashiCorp AzureAD Terraform Provider](https://github.com/hashicorp/terraform-provider-azuread).

## Contributing

This Microsoft Graph SDK is auto-generated using the [Pandora project](https://github.com/hashicorp/pandora), which obtains API definitions from [OpenAPI3 specs published by Microsoft](https://github.com/microsoftgraph/msgraph-metadata).

To add new services to this SDK, please amend the [`microsoft-graph.hcl`](https://github.com/hashicorp/pandora/tree/main/config) configuration file in the hashicorp/pandora repository, and open a pull request.
