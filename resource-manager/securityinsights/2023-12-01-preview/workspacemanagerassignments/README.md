
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerassignments` Documentation

The `workspacemanagerassignments` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/workspacemanagerassignments"
```


### Client Initialization

```go
client := workspacemanagerassignments.NewWorkspaceManagerAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceManagerAssignmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceManagerAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName")

payload := workspacemanagerassignments.WorkspaceManagerAssignment{
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


### Example Usage: `WorkspaceManagerAssignmentsClient.Delete`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceManagerAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.Get`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceManagerAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.List`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, workspacemanagerassignments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workspacemanagerassignments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.WorkspaceManagerAssignmentJobsCreate`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceManagerAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName")

read, err := client.WorkspaceManagerAssignmentJobsCreate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.WorkspaceManagerAssignmentJobsDelete`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName", "jobName")

read, err := client.WorkspaceManagerAssignmentJobsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.WorkspaceManagerAssignmentJobsGet`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName", "jobName")

read, err := client.WorkspaceManagerAssignmentJobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceManagerAssignmentsClient.WorkspaceManagerAssignmentJobsList`

```go
ctx := context.TODO()
id := workspacemanagerassignments.NewWorkspaceManagerAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "workspaceManagerAssignmentName")

// alternatively `client.WorkspaceManagerAssignmentJobsList(ctx, id, workspacemanagerassignments.DefaultWorkspaceManagerAssignmentJobsListOperationOptions())` can be used to do batched pagination
items, err := client.WorkspaceManagerAssignmentJobsListComplete(ctx, id, workspacemanagerassignments.DefaultWorkspaceManagerAssignmentJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
