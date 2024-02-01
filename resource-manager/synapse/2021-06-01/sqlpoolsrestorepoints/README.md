
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsrestorepoints` Documentation

The `sqlpoolsrestorepoints` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsrestorepoints"
```


### Client Initialization

```go
client := sqlpoolsrestorepoints.NewSqlPoolsRestorePointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsRestorePointsClient.SqlPoolRestorePointsCreate`

```go
ctx := context.TODO()
id := sqlpoolsrestorepoints.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolsrestorepoints.CreateSqlPoolRestorePointDefinition{
	// ...
}


if err := client.SqlPoolRestorePointsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SqlPoolsRestorePointsClient.SqlPoolRestorePointsDelete`

```go
ctx := context.TODO()
id := sqlpoolsrestorepoints.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "restorePointValue")

read, err := client.SqlPoolRestorePointsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsRestorePointsClient.SqlPoolRestorePointsGet`

```go
ctx := context.TODO()
id := sqlpoolsrestorepoints.NewRestorePointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue", "restorePointValue")

read, err := client.SqlPoolRestorePointsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
