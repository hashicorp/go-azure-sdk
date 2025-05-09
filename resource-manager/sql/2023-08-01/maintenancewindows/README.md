
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/maintenancewindows` Documentation

The `maintenancewindows` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/maintenancewindows"
```


### Client Initialization

```go
client := maintenancewindows.NewMaintenanceWindowsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MaintenanceWindowsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

payload := maintenancewindows.MaintenanceWindows{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, maintenancewindows.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MaintenanceWindowsClient.Get`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

read, err := client.Get(ctx, id, maintenancewindows.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
