
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionrepetitiondefinitions` Documentation

The `workflowrunactionrepetitiondefinitions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionrepetitiondefinitions"
```


### Client Initialization

```go
client := workflowrunactionrepetitiondefinitions.NewWorkflowRunActionRepetitionDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunActionRepetitionDefinitionsClient.WorkflowRunActionRepetitionsGet`

```go
ctx := context.TODO()
id := workflowrunactionrepetitiondefinitions.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName", "repetitionName")

read, err := client.WorkflowRunActionRepetitionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionRepetitionDefinitionsClient.WorkflowRunActionRepetitionsList`

```go
ctx := context.TODO()
id := workflowrunactionrepetitiondefinitions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName")

// alternatively `client.WorkflowRunActionRepetitionsList(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionRepetitionDefinitionsClient.WorkflowRunActionRepetitionsListExpressionTraces`

```go
ctx := context.TODO()
id := workflowrunactionrepetitiondefinitions.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName", "repetitionName")

// alternatively `client.WorkflowRunActionRepetitionsListExpressionTraces(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsListExpressionTracesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
