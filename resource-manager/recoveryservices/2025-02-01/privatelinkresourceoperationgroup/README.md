
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/privatelinkresourceoperationgroup` Documentation

The `privatelinkresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/privatelinkresourceoperationgroup"
```


### Client Initialization

```go
client := privatelinkresourceoperationgroup.NewPrivateLinkResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourceOperationGroupClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := privatelinkresourceoperationgroup.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "privateLinkResourceName")

read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourceOperationGroupClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := privatelinkresourceoperationgroup.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.PrivateLinkResourcesList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkResourcesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
