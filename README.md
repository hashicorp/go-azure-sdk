# `hashicorp/go-azure-sdk`

This SDK is an opinionated Go SDK for Azure Resource Manager, primarily intended to be used in the [Terraform AzureRM Provider](https://github.com/hashicorp/terraform-provider-azurerm) and the [Packer Azure Plugin](https://github.com/hashicorp/packer-plugin-azure).

The SDKs within this repository are generated using the data in [the Azure Rest API Specifications repository](https://github.com/Azure/azure-rest-api-specs).

At this time this SDK uses both [the base-layer within this repository](./sdk) and [the `Azure/go-autorest` base layer](https://github.com/Azure/go-autorest) (although this is gradually being removed).

## Repository Structure

This repository contains:

* `./docs` - usage documentation.
* `./resource-manager` - the Resource Manager Services supported by this SDK.
* `./sdk` - the base layer that's used in this repository.

In addition, a `README.md` file can be found within each Service / API Version in the `./resource-manager` directory highlighting how that SDK can be used ([example](resource-manager/aadb2c/2021-04-01-preview/tenants)).

## Getting Started

SDKs within this repository can be found within the `./resource-manager` directory, in the structure `./resource-manager/{serviceName}/{apiVersion}/{resource}` for example [this the Resource Groups SDK for API Version `2022-09-01` can be found in `./resource-manager/resources/2022-09-01/resourcegroups`](./resource-manager/resources/2022-09-01/resourcegroups), where the README for which also demonstrates how to use each function.

SDK Clients require authentication which is available from the `auth` package, see [the README in the `./sdk/auth` directory](./sdk/auth) for information on how to authenticate and obtain an Authorizer.

Once you've got an Authorizer, you can use an SDK Client like so:

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-09-01/resourcegroups"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5 * time.Minute)
	defer cancel()

	id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")
	environment := environments.AzurePublic()
	credentials := auth.Credentials{
		Environment:                       *environment,
		EnableAuthenticatingUsingAzureCLI: true,
		// you can also authenticate using a Service Principal, via MSI, OIDC etc by turning on the flag
	}
	log.Printf("[DEBUG] Obtaining Credentials..")
	authorizer, err := auth.NewAuthorizerFromCredentials(ctx, credentials, environment.ResourceManager)
	if err != nil {
		log.Fatalf("building authorizer from credentials: %+v", err)
	}

	log.Printf("[DEBUG] Building Resource Groups Client..")
	client, err := resourcegroups.NewResourceGroupsClientWithBaseURI(environment.ResourceManager)
	if err != nil {
		log.Fatalf("building ResourceGroups client: %+v", err)
	}
	client.Client.Authorizer = authorizer

	// NOTE: the Resource ID type can be used for logging too, it contains a human-friendly description
	log.Printf("[DEBUG] Creating %s..", id)
	payload := resourcegroups.ResourceGroup{
		// ...
	}
	if _, err := client.CreateOrUpdate(ctx, id, payload); err != nil {
		log.Fatalf("creating %s: %+v", id, err)
	}
	log.Printf("[DEBUG] %s created..", id)
}
```

## Further Documentation

Further information [can be found in the `./docs` directory](./docs), and also the `README.md` file within each Service / API Version in the `./resource-manager` directory ([example](resource-manager/aadb2c/2021-04-01-preview/tenants)).
