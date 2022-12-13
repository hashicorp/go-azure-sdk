
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2021-12-01/listprivatelinkresources` Documentation

The `listprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `recoveryservices` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2021-12-01/listprivatelinkresources"
```


### Client Initialization

```go
client := listprivatelinkresources.NewListPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListPrivateLinkResourcesClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := listprivatelinkresources.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.PrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
