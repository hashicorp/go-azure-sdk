
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/databasetables` Documentation

The `databasetables` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/databasetables"
```


### Client Initialization

```go
client := databasetables.NewDatabaseTablesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseTablesClient.Get`

```go
ctx := context.TODO()
id := databasetables.NewTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue", "tableValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseTablesClient.ListBySchema`

```go
ctx := context.TODO()
id := databasetables.NewSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "schemaValue")

// alternatively `client.ListBySchema(ctx, id, databasetables.DefaultListBySchemaOperationOptions())` can be used to do batched pagination
items, err := client.ListBySchemaComplete(ctx, id, databasetables.DefaultListBySchemaOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
