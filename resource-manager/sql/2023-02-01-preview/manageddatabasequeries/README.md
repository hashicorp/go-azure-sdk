
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabasequeries` Documentation

The `manageddatabasequeries` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabasequeries"
```


### Client Initialization

```go
client := manageddatabasequeries.NewManagedDatabaseQueriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseQueriesClient.Get`

```go
ctx := context.TODO()
id := manageddatabasequeries.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "queryId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDatabaseQueriesClient.ListByQuery`

```go
ctx := context.TODO()
id := manageddatabasequeries.NewQueryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName", "queryId")

// alternatively `client.ListByQuery(ctx, id, manageddatabasequeries.DefaultListByQueryOperationOptions())` can be used to do batched pagination
items, err := client.ListByQueryComplete(ctx, id, manageddatabasequeries.DefaultListByQueryOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
