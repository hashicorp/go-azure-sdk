
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnectionoperationgroup` Documentation

The `databaseconnectionoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnectionoperationgroup"
```


### Client Initialization

```go
client := databaseconnectionoperationgroup.NewDatabaseConnectionOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesCreateOrUpdateDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "databaseConnectionName")

payload := databaseconnectionoperationgroup.DatabaseConnection{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateDatabaseConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesDeleteDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "databaseConnectionName")

read, err := client.StaticSitesDeleteDatabaseConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesGetDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "databaseConnectionName")

read, err := client.StaticSitesGetDatabaseConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesGetDatabaseConnectionWithDetails`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "databaseConnectionName")

read, err := client.StaticSitesGetDatabaseConnectionWithDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesGetDatabaseConnections`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewStaticSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName")

// alternatively `client.StaticSitesGetDatabaseConnections(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetDatabaseConnectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatabaseConnectionOperationGroupClient.StaticSitesUpdateDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnectionoperationgroup.NewDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "databaseConnectionName")

payload := databaseconnectionoperationgroup.DatabaseConnectionPatchRequest{
	// ...
}


read, err := client.StaticSitesUpdateDatabaseConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
