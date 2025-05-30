
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasecolumns` Documentation

The `manageddatabasecolumns` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasecolumns"
```


### Client Initialization

```go
client := manageddatabasecolumns.NewManagedDatabaseColumnsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseColumnsClient.Get`

```go
ctx := context.TODO()
id := manageddatabasecolumns.NewTableColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName", "columnName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseColumnsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id, manageddatabasecolumns.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, manageddatabasecolumns.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDatabaseColumnsClient.ListByTable`

```go
ctx := context.TODO()
id := manageddatabasecolumns.NewSchemaTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "schemaName", "tableName")

// alternatively `client.ListByTable(ctx, id, manageddatabasecolumns.DefaultListByTableOperationOptions())` can be used to do batched pagination
items, err := client.ListByTableComplete(ctx, id, manageddatabasecolumns.DefaultListByTableOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
