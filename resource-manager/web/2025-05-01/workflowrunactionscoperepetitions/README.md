
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionscoperepetitions` Documentation

The `workflowrunactionscoperepetitions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/workflowrunactionscoperepetitions"
```


### Client Initialization

```go
client := workflowrunactionscoperepetitions.NewWorkflowRunActionScopeRepetitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunActionScopeRepetitionsClient.Get`

```go
ctx := context.TODO()
id := workflowrunactionscoperepetitions.NewScopeRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName", "scopeRepetitionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionScopeRepetitionsClient.List`

```go
ctx := context.TODO()
id := workflowrunactionscoperepetitions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
