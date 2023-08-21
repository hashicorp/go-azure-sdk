# `hashicorp/go-azure-sdk` Documentation

## Overview

* [Authentication (from the `./sdk/auth` directory)](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/auth/README.md).
* [Working with Discriminated Types](working-with-discriminators.md).
* [Frequently Asked Question](frequently-asked-questions.md).

Additional README's can be found within each the directory for each Service / API Version ([example](https://github.com/hashicorp/go-azure-sdk/blob/main/resource-manager/resources/2022-09-01/resourcegroups/README.md)).

## Example Usage

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
