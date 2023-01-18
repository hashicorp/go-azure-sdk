
## `github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowruns` Documentation

The `workflowruns` SDK allows for interaction with the Azure Resource Manager Service `logic` (API Version `2019-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowruns"
```


### Client Initialization

```go
client := workflowruns.NewWorkflowRunsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunsClient.Cancel`

```go
ctx := context.TODO()
id := workflowruns.NewRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workflowValue", "runValue")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunsClient.Get`

```go
ctx := context.TODO()
id := workflowruns.NewRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workflowValue", "runValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunsClient.List`

```go
ctx := context.TODO()
id := workflowruns.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workflowValue")

// alternatively `client.List(ctx, id, workflowruns.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowruns.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
