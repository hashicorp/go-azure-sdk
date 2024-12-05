
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagergroups` Documentation

The `workspacemanagergroups` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagergroups"
```


### Client Initialization

```go
client := workspacemanagergroups.NewWorkspaceManagerGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagerGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagergroups.NewWorkspaceManagerGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerGroupName")

payload := workspacemanagergroups.WorkspaceManagerGroup{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerGroupsClient.Delete`

```go
ctx := context.TODO()
id := workspacemanagergroups.NewWorkspaceManagerGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerGroupName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerGroupsClient.Get`

```go
ctx := context.TODO()
id := workspacemanagergroups.NewWorkspaceManagerGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerGroupsClient.List`

```go
ctx := context.TODO()
id := workspacemanagergroups.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, workspacemanagergroups.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workspacemanagergroups.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
