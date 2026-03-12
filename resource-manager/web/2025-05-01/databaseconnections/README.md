
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnections` Documentation

The `databaseconnections` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/databaseconnections"
```


### Client Initialization

```go
client := databaseconnections.NewDatabaseConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesCreateOrUpdateBuildDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "databaseConnectionName")

payload := databaseconnections.DatabaseConnection{
	// ...
}


read, err := client.StaticSitesCreateOrUpdateBuildDatabaseConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesDeleteBuildDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "databaseConnectionName")

read, err := client.StaticSitesDeleteBuildDatabaseConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesGetBuildDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "databaseConnectionName")

read, err := client.StaticSitesGetBuildDatabaseConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesGetBuildDatabaseConnectionWithDetails`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "databaseConnectionName")

read, err := client.StaticSitesGetBuildDatabaseConnectionWithDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesGetBuildDatabaseConnections`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName")

// alternatively `client.StaticSitesGetBuildDatabaseConnections(ctx, id)` can be used to do batched pagination
items, err := client.StaticSitesGetBuildDatabaseConnectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatabaseConnectionsClient.StaticSitesUpdateBuildDatabaseConnection`

```go
ctx := context.TODO()
id := databaseconnections.NewBuildDatabaseConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "staticSiteName", "buildName", "databaseConnectionName")

payload := databaseconnections.DatabaseConnectionPatchRequest{
	// ...
}


read, err := client.StaticSitesUpdateBuildDatabaseConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
