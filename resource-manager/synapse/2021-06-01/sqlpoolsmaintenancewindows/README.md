
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsmaintenancewindows` Documentation

The `sqlpoolsmaintenancewindows` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsmaintenancewindows"
```


### Client Initialization

```go
client := sqlpoolsmaintenancewindows.NewSqlPoolsMaintenanceWindowsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsMaintenanceWindowsClient.SqlPoolMaintenanceWindowsCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolsmaintenancewindows.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolsmaintenancewindows.MaintenanceWindows{
	// ...
}


read, err := client.SqlPoolMaintenanceWindowsCreateOrUpdate(ctx, id, payload, sqlpoolsmaintenancewindows.DefaultSqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsMaintenanceWindowsClient.SqlPoolMaintenanceWindowsGet`

```go
ctx := context.TODO()
id := sqlpoolsmaintenancewindows.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

read, err := client.SqlPoolMaintenanceWindowsGet(ctx, id, sqlpoolsmaintenancewindows.DefaultSqlPoolMaintenanceWindowsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
