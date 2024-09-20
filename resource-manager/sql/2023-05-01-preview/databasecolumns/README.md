
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/databasecolumns` Documentation

The `databasecolumns` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/databasecolumns"
```


### Client Initialization

```go
client := databasecolumns.NewDatabaseColumnsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseColumnsClient.Get`

```go
ctx := context.TODO()
id := databasecolumns.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "schemaName", "tableName", "columnName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseColumnsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id, databasecolumns.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, databasecolumns.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatabaseColumnsClient.ListByTable`

```go
ctx := context.TODO()
id := databasecolumns.NewTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "schemaName", "tableName")

// alternatively `client.ListByTable(ctx, id, databasecolumns.DefaultListByTableOperationOptions())` can be used to do batched pagination
items, err := client.ListByTableComplete(ctx, id, databasecolumns.DefaultListByTableOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
