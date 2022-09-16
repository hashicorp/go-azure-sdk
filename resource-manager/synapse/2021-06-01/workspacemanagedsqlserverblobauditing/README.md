
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverblobauditing` Documentation

The `workspacemanagedsqlserverblobauditing` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverblobauditing"
```


### Client Initialization

```go
client := workspacemanagedsqlserverblobauditing.NewWorkspaceManagedSqlServerBlobAuditingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspacemanagedsqlserverblobauditing.ServerBlobAuditingPolicy{
	// ...
}


if err := client.WorkspaceManagedSqlServerBlobAuditingPoliciesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerBlobAuditingPoliciesGet`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.WorkspaceManagedSqlServerBlobAuditingPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerBlobAuditingPoliciesListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspacemanagedsqlserverblobauditing.ExtendedServerBlobAuditingPolicy{
	// ...
}


if err := client.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerBlobAuditingClient.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace`

```go
ctx := context.TODO()
id := workspacemanagedsqlserverblobauditing.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerExtendedBlobAuditingPoliciesListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
