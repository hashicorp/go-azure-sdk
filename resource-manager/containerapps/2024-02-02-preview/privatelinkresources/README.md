
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.ManagedEnvironmentPrivateLinkResourcesList`

```go
ctx := context.TODO()
id := privatelinkresources.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName")

// alternatively `client.ManagedEnvironmentPrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentPrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
