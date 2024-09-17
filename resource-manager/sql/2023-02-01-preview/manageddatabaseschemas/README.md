
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabaseschemas` Documentation

The `manageddatabaseschemas` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabaseschemas"
```


### Client Initialization

```go
client := manageddatabaseschemas.NewManagedDatabaseSchemasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseSchemasClient.Get`

```go
ctx := context.TODO()
id := manageddatabaseschemas.NewDatabaseSchemaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue", "schemaValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseSchemasClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

// alternatively `client.ListByDatabase(ctx, id, manageddatabaseschemas.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, manageddatabaseschemas.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
