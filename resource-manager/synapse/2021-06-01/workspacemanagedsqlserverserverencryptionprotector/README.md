
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverserverencryptionprotector` Documentation

The `workspacemanagedsqlserverserverencryptionprotector` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverserverencryptionprotector"
```


### Client Initialization

```go
client := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceManagedSqlServerServerEncryptionProtectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerServerEncryptionProtectorClient.WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspacemanagedsqlserverserverencryptionprotector.EncryptionProtector{
	// ...
}


if err := client.WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspaceManagedSqlServerServerEncryptionProtectorClient.WorkspaceManagedSqlServerEncryptionProtectorGet`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.WorkspaceManagedSqlServerEncryptionProtectorGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerServerEncryptionProtectorClient.WorkspaceManagedSqlServerEncryptionProtectorList`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerEncryptionProtectorList(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerEncryptionProtectorListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceManagedSqlServerServerEncryptionProtectorClient.WorkspaceManagedSqlServerEncryptionProtectorRevalidate`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

if err := client.WorkspaceManagedSqlServerEncryptionProtectorRevalidateThenPoll(ctx, id); err != nil {
	// handle the error
}
```
