
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/restorabledroppedsqlpools` Documentation

The `restorabledroppedsqlpools` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/restorabledroppedsqlpools"
```


### Client Initialization

```go
client := restorabledroppedsqlpools.NewRestorableDroppedSqlPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestorableDroppedSqlPoolsClient.Get`

```go
ctx := context.TODO()
id := restorabledroppedsqlpools.NewRestorableDroppedSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "restorableDroppedSqlPoolId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorableDroppedSqlPoolsClient.ListByWorkspace`

```go
ctx := context.TODO()
id := restorabledroppedsqlpools.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.ListByWorkspace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
