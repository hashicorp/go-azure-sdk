
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/workflowtriggerhistories` Documentation

The `workflowtriggerhistories` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/workflowtriggerhistories"
```


### Client Initialization

```go
client := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowTriggerHistoriesClient.Get`

```go
ctx := context.TODO()
id := workflowtriggerhistories.NewTriggerHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "triggerName", "historyName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowTriggerHistoriesClient.List`

```go
ctx := context.TODO()
id := workflowtriggerhistories.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "triggerName")

// alternatively `client.List(ctx, id, workflowtriggerhistories.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowtriggerhistories.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowTriggerHistoriesClient.Resubmit`

```go
ctx := context.TODO()
id := workflowtriggerhistories.NewTriggerHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "triggerName", "historyName")

if err := client.ResubmitThenPoll(ctx, id); err != nil {
	// handle the error
}
```
