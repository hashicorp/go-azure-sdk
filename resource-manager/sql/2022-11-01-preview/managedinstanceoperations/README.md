
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/managedinstanceoperations` Documentation

The `managedinstanceoperations` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/managedinstanceoperations"
```


### Client Initialization

```go
client := managedinstanceoperations.NewManagedInstanceOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstanceOperationsClient.Cancel`

```go
ctx := context.TODO()
id := managedinstanceoperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "operationIdValue")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedInstanceOperationsClient.Get`

```go
ctx := context.TODO()
id := managedinstanceoperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "operationIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedInstanceOperationsClient.ListByManagedInstance`

```go
ctx := context.TODO()
id := managedinstanceoperations.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByManagedInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
