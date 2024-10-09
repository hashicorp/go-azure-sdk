
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserver` Documentation

The `workspacemanagedsqlserver` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserver"
```


### Client Initialization

```go
client := workspacemanagedsqlserver.NewWorkspaceManagedSqlServerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerClient.RecoverableSqlPoolsGet`

```go
ctx := context.TODO()
id := workspacemanagedsqlserver.NewRecoverableSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "recoverableSqlPoolName")

read, err := client.RecoverableSqlPoolsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerClient.RecoverableSqlPoolsList`

```go
ctx := context.TODO()
id := workspacemanagedsqlserver.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.RecoverableSqlPoolsList(ctx, id)` can be used to do batched pagination
items, err := client.RecoverableSqlPoolsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
