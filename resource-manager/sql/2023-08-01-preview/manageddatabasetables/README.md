
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/manageddatabasetables` Documentation

The `manageddatabasetables` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/manageddatabasetables"
```


### Client Initialization

```go
client := manageddatabasetables.NewManagedDatabaseTablesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseTablesClient.Get`

```go
ctx := context.TODO()
id := manageddatabasetables.NewSchemaTableID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue", "tableValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseTablesClient.ListBySchema`

```go
ctx := context.TODO()
id := manageddatabasetables.NewDatabaseSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue")

// alternatively `client.ListBySchema(ctx, id, manageddatabasetables.DefaultListBySchemaOperationOptions())` can be used to do batched pagination
items, err := client.ListBySchemaComplete(ctx, id, manageddatabasetables.DefaultListBySchemaOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
