
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/manageddatabasemoveoperations` Documentation

The `manageddatabasemoveoperations` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/manageddatabasemoveoperations"
```


### Client Initialization

```go
client := manageddatabasemoveoperations.NewManagedDatabaseMoveOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseMoveOperationsClient.ListByLocation`

```go
ctx := context.TODO()
id := manageddatabasemoveoperations.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

// alternatively `client.ListByLocation(ctx, id, manageddatabasemoveoperations.DefaultListByLocationOperationOptions())` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id, manageddatabasemoveoperations.DefaultListByLocationOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
