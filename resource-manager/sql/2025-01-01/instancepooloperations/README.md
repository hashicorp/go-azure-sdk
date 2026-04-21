
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2025-01-01/instancepooloperations` Documentation

The `instancepooloperations` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2025-01-01/instancepooloperations"
```


### Client Initialization

```go
client := instancepooloperations.NewInstancePoolOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstancePoolOperationsClient.Get`

```go
ctx := context.TODO()
id := instancepooloperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolName", "operationId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstancePoolOperationsClient.ListByInstancePool`

```go
ctx := context.TODO()
id := instancepooloperations.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolName")

// alternatively `client.ListByInstancePool(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstancePoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
