# `hashicorp/go-azure-sdk` Documentation have error

## Overview

* [Authentication error](authentication.md).
* [Working with Discriminated Types](working-with-discriminators.md).
* [Frequently Asked Question](frequently-asked-questions.md).

Additional README's can be found within each the directory for each Service / API Version ([example](XXX)).

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resourcegroups"
	"github.com/manicminer/hamilton/environments"
)

func main() {
	// resourceGroupName is the name of the Resource Group to create
	resourceGroupName := "go-azure-sdk-resources"

	auth := authentication.Builder{
		ClientID:                 os.Getenv("ARM_CLIENT_ID"),
		ClientSecret:             os.Getenv("ARM_CLIENT_SECRET"),
		SubscriptionID:           os.Getenv("ARM_SUBSCRIPTION_ID"),
		TenantID:                 os.Getenv("ARM_TENANT_ID"),
		Environment:              "public",
		SupportsClientSecretAuth: true,
		UseMicrosoftGraph:        true,

		// other authentication methods can be enabled using:
		//SupportsAzureCliToken:          false,
		//SupportsManagedServiceIdentity: false,
		//SupportsClientCertAuth:         false,
	}
	if err := run(context.TODO(), auth, resourceGroupName); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

func run(ctx context.Context, builder authentication.Builder, resourceGroupName string) error {
	// first load the Environment (containing metadata about the Azure Environment)
	env, err := environments.EnvironmentFromString(builder.Environment)
	if err != nil {
		return fmt.Errorf("locating environment %q: %+v", builder.Environment, err)
	}

	// then obtain a Go-AutoRest Authorizer using the credentials
	auth, err := buildAuthorizer(ctx, builder, env)
	if err != nil {
		return fmt.Errorf("building authorizer: %+v", err)
	}

	// then create a Resource Groups Client, targeting Azure Public
	client := resourcegroups.NewResourceGroupsClientWithBaseURI(string(env.ResourceManager.Endpoint))

	// add authorization to this Client
	client.Client.Authorizer = auth

	// define the Resource ID for this Resource Group
	id := commonids.NewResourceGroupID(builder.SubscriptionID, resourceGroupName)

	// build the payload
	payload := resourcegroups.ResourceGroup{
		Location: "westeurope",
		Tags: &map[string]string{
			"hello": "world",
		},
	}

	// Create the Resource Group
	log.Printf("[DEBUG] Creating %s..", id)
	if _, err := client.CreateOrUpdate(ctx, id, payload); err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}
	log.Printf("[DEBUG] Created %s (ID %q).", id, id.ID())
	return nil
}

func buildAuthorizer(ctx context.Context, builder authentication.Builder, env environments.Environment) (autorest.Authorizer, error) {
	config, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("building auth credentials: %+v", err)
	}

	authorizer, err := config.GetMSALToken(ctx, env.ResourceManager, nil, nil, string(env.ResourceManager.Endpoint))
	if err != nil {
		return nil, fmt.Errorf("retrieving MSAL Token: %+v", err)
	}
	return authorizer, nil
}
```
