
## `github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listserviceproviders` Documentation

The `listserviceproviders` SDK allows for interaction with the Azure Resource Manager Service `botservice` (API Version `2022-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listserviceproviders"
```


### Client Initialization

```go
client := listserviceproviders.NewListServiceProvidersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListServiceProvidersClient.BotConnectionListServiceProviders`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.BotConnectionListServiceProviders(ctx, id)` can be used to do batched pagination
items, err := client.BotConnectionListServiceProvidersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
