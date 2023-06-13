
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/privatelinkresource` Documentation

The `privatelinkresource` SDK allows for interaction with the Azure Resource Manager Service `batch` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/privatelinkresource"
```


### Client Initialization

```go
client := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourceClient.Get`

```go
ctx := context.TODO()
id := privatelinkresource.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "privateLinkResourceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourceClient.ListByBatchAccount`

```go
ctx := context.TODO()
id := privatelinkresource.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue")

// alternatively `client.ListByBatchAccount(ctx, id, privatelinkresource.DefaultListByBatchAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBatchAccountComplete(ctx, id, privatelinkresource.DefaultListByBatchAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
