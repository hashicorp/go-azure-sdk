
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversqlusages` Documentation

The `workspacemanagedsqlserversqlusages` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversqlusages"
```


### Client Initialization

```go
client := workspacemanagedsqlserversqlusages.NewWorkspaceManagedSqlServerSqlUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerSqlUsagesClient.WorkspaceManagedSqlServerUsagesList`

```go
ctx := context.TODO()
id := workspacemanagedsqlserversqlusages.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
