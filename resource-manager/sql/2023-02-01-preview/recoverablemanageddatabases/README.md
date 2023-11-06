
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/recoverablemanageddatabases` Documentation

The `recoverablemanageddatabases` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/recoverablemanageddatabases"
```


### Client Initialization

```go
client := recoverablemanageddatabases.NewRecoverableManagedDatabasesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoverableManagedDatabasesClient.Get`

```go
ctx := context.TODO()
id := recoverablemanageddatabases.NewManagedInstanceRecoverableDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "recoverableDatabaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoverableManagedDatabasesClient.ListByInstance`

```go
ctx := context.TODO()
id := recoverablemanageddatabases.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
