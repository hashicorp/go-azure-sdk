
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabaserestoredetails` Documentation

The `manageddatabaserestoredetails` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-02-01-preview/manageddatabaserestoredetails"
```


### Client Initialization

```go
client := manageddatabaserestoredetails.NewManagedDatabaseRestoreDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseRestoreDetailsClient.Get`

```go
ctx := context.TODO()
id := manageddatabaserestoredetails.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "databaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
