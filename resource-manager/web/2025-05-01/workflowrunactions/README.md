
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactions` Documentation

The `workflowrunactions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactions"
```


### Client Initialization

```go
client := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunActionsClient.Get`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionsClient.List`

```go
ctx := context.TODO()
id := workflowrunactions.NewRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName")

// alternatively `client.List(ctx, id, workflowrunactions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowrunactions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.ListExpressionTraces`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName")

// alternatively `client.ListExpressionTraces(ctx, id)` can be used to do batched pagination
items, err := client.ListExpressionTracesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
