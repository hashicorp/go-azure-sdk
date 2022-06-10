# `hashicorp/go-azure-sdk`

This SDK is an opinionated Go SDK for Azure Resource Manager, primarily intended to be used [in the Terraform AzureRM Provider](https://github.com/hashicorp/terraform-provider-azurerm).

The SDKs within this repository are generated using the data in [the Azure Rest API Specifications repository](https://github.com/Azure/azure-rest-api-specs).

At this time this SDK makes use of [the `Azure/go-autorest` base layer](https://github.com/Azure/go-autorest), however this will change in a future release.

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2020-05-01/signalr"
)

func main() {
	auth := authConfig{
		clientId:       os.Getenv("ARM_CLIENT_ID"),
		clientSecret:   os.Getenv("ARM_CLIENT_SECRET"),
		subscriptionId: os.Getenv("ARM_SUBSCRIPTION_ID"),
		tenantId:       os.Getenv("ARM_TENANT_ID"),
		environment:    azure.PublicCloud,
	}
	authorizer, err := auth.buildAuthorizer()
	if err != nil {
		log.Fatalf("building authorizer: %+v", err)
	}

	signalRClient := signalr.NewSignalRClientWithBaseURI(auth.environment.ResourceManagerEndpoint)
	signalRClient.Client.Authorizer = authorizer

	id := signalr.NewSignalRID(auth.subscriptionId, "my-resource-group", "an-existing-signalr")
	log.Printf("[DEBUG] Retrieving %s", id)
	read, err := signalRClient.Get(context.TODO(), id)
	if err != nil {
		log.Fatalf("retrieving %s: %+v", id, err)
	}
	if response.WasNotFound(read.HttpResponse) {
		log.Fatalf("%s was not found", id)
	}
	if model := read.Model; model != nil {
		log.Printf("The name is %s", *model.Name)
	}
}

type authConfig struct {
	clientId       string
	clientSecret   string
	subscriptionId string
	tenantId       string
	environment    azure.Environment
}

func (c authConfig) buildAuthorizer() (autorest.Authorizer, error) {
	oauth, err := adal.NewOAuthConfig(c.environment.ActiveDirectoryEndpoint, c.tenantId)
	if err != nil {
		return nil, fmt.Errorf("building oauth config: %+v", err)
	}

	spt, err := adal.NewServicePrincipalToken(*oauth, c.clientId, c.clientSecret, c.environment.ResourceManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("building access token: %+v", err)
	}

	auth := autorest.NewBearerAuthorizer(spt)
	return auth, nil
}
```

