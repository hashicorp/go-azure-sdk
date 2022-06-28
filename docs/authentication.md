# Authentication

At this time this SDK makes use of the `Azure/go-autorest` base layer - which uses [this interface for authentication](https://github.com/Azure/go-autorest/blob/7dd32b67be4e6c9386b9ba7b1c44a51263f05270/autorest/authorization.go#L42-L44).

Whilst the `Azure/go-autorest` base layer has native support for authentication, this makes use of ADAL (and Azure Active Directory) which is now deprecated - and the recommendation is to use MSAL (and Microsoft Graph).

In the future this SDK will be updated to use a different base layer which has native support for MSAL - however in the interim we can use [the MSAL support within the Hamilton library](https://github.com/manicminer/hamilton) to authenticate.

## Example: Authenticating using MSAL (and Microsoft Graph)

Support for authenticating using MSAL (and Microsoft Graph) can be found in [the Hamilton library](https://github.com/manicminer/hamilton) - and can be used as below:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resourcegroups"
	"github.com/manicminer/hamilton/environments"
)

func main() {
	auth, err := buildAuthorizer(context.TODO())
	if err != nil {
		log.Fatalf("building authorizer: %+v", err)
	}

	// then create a Resource Groups Client, targeting Azure Public
	client := resourcegroups.NewResourceGroupsClientWithBaseURI(string(environments.ResourceManagerPublicEndpoint))

	// add authorization to this Client
	client.Client.Authorizer = auth

	// ...
}

func buildAuthorizer(ctx context.Context) (autorest.Authorizer, error) {
	// create a builder, which determines which authentication method should be used based on the
	// credentials provided/toggles enabled (e.g. Client Secret, Client Certificate, MSI etc).
	builder := authentication.Builder{
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

	// then load the Environment (containing metadata about the Azure Environment)
	env, err := environments.EnvironmentFromString(builder.Environment)
	if err != nil {
		return nil, fmt.Errorf("locating environment %q: %+v", builder.Environment, err)
	}

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
