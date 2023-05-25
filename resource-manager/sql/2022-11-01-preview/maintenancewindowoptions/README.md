
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/maintenancewindowoptions` Documentation

The `maintenancewindowoptions` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2022-11-01-preview/maintenancewindowoptions"
```


### Client Initialization

```go
client := maintenancewindowoptions.NewMaintenanceWindowOptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MaintenanceWindowOptionsClient.Get`

```go
ctx := context.TODO()
id := maintenancewindowoptions.NewDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue")

read, err := client.Get(ctx, id, maintenancewindowoptions.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
