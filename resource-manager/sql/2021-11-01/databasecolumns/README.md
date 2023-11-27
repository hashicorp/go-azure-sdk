
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/databasecolumns` Documentation

The `databasecolumns` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/databasecolumns"
```


### Client Initialization

```go
client := databasecolumns.NewDatabaseColumnsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseColumnsClient.Get`

```go
ctx := context.TODO()
id := databasecolumns.NewColumnID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue", "columnValue")

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
id := databasecolumns.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

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
id := databasecolumns.NewTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue")

// alternatively `client.ListByTable(ctx, id, databasecolumns.DefaultListByTableOperationOptions())` can be used to do batched pagination
items, err := client.ListByTableComplete(ctx, id, databasecolumns.DefaultListByTableOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
