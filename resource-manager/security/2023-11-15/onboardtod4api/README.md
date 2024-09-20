
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/onboardtod4api` Documentation

The `onboardtod4api` SDK allows for interaction with Azure Resource Manager `security` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/onboardtod4api"
```


### Client Initialization

```go
client := onboardtod4api.NewOnboardToD4APIClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnboardToD4APIClient.APICollectionsOnboardAzureApiManagementApi`

```go
ctx := context.TODO()
id := onboardtod4api.NewApiCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceName", "apiId")

if err := client.APICollectionsOnboardAzureApiManagementApiThenPoll(ctx, id); err != nil {
	// handle the error
}
```
