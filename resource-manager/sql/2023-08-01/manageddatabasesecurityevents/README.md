
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasesecurityevents` Documentation

The `manageddatabasesecurityevents` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/manageddatabasesecurityevents"
```


### Client Initialization

```go
client := manageddatabasesecurityevents.NewManagedDatabaseSecurityEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDatabaseSecurityEventsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id, manageddatabasesecurityevents.DefaultListByDatabaseOperationOptions())` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id, manageddatabasesecurityevents.DefaultListByDatabaseOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
