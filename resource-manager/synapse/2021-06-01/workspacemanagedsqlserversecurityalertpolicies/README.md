
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversecurityalertpolicies` Documentation

The `workspacemanagedsqlserversecurityalertpolicies` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversecurityalertpolicies"
```


### Client Initialization

```go
client := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceManagedSqlServerSecurityAlertPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagedSqlServerSecurityAlertPoliciesClient.WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspacemanagedsqlserversecurityalertpolicies.ServerSecurityAlertPolicy{
	// ...
}


if err := client.WorkspaceManagedSqlServerSecurityAlertPolicyCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspaceManagedSqlServerSecurityAlertPoliciesClient.WorkspaceManagedSqlServerSecurityAlertPolicyGet`

```go
ctx := context.TODO()
id := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.WorkspaceManagedSqlServerSecurityAlertPolicyGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagedSqlServerSecurityAlertPoliciesClient.WorkspaceManagedSqlServerSecurityAlertPolicyList`

```go
ctx := context.TODO()
id := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.WorkspaceManagedSqlServerSecurityAlertPolicyList(ctx, id)` can be used to do batched pagination
items, err := client.WorkspaceManagedSqlServerSecurityAlertPolicyListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
