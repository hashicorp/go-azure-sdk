
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/recoverablesqlpools` Documentation

The `recoverablesqlpools` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/recoverablesqlpools"
```


### Client Initialization

```go
client := recoverablesqlpools.NewRecoverableSqlPoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoverableSqlPoolsClient.WorkspaceManagedSqlServerRecoverableSqlPoolsGet`

```go
ctx := context.TODO()
id := recoverablesqlpools.NewRecoverableSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

read, err := client.WorkspaceManagedSqlServerRecoverableSqlPoolsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoverableSqlPoolsClient.WorkspaceManagedSqlServerRecoverableSqlPoolsList`

```go
ctx := context.TODO()
id := recoverablesqlpools.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerRecoverableSqlPoolsList(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerRecoverableSqlPoolsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
