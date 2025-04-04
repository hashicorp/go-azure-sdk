
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagermember` Documentation

The `workspacemanagermember` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagermember"
```


### Client Initialization

```go
client := workspacemanagermember.NewWorkspaceManagerMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagerMemberClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagermember.NewWorkspaceManagerMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerMemberName")

payload := workspacemanagermember.WorkspaceManagerMember{
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


### Example Usage: `WorkspaceManagerMemberClient.Delete`

```go
ctx := context.TODO()
id := workspacemanagermember.NewWorkspaceManagerMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerMemberName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerMemberClient.Get`

```go
ctx := context.TODO()
id := workspacemanagermember.NewWorkspaceManagerMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerMemberName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerMemberClient.List`

```go
ctx := context.TODO()
id := workspacemanagermember.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, workspacemanagermember.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workspacemanagermember.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
