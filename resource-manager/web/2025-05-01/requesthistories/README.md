
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/requesthistories` Documentation

The `requesthistories` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/requesthistories"
```


### Client Initialization

```go
client := requesthistories.NewRequestHistoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RequestHistoriesClient.WorkflowRunActionRepetitionsRequestHistoriesGet`

```go
ctx := context.TODO()
id := requesthistories.NewRequestHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName", "repetitionName", "requestHistoryName")

read, err := client.WorkflowRunActionRepetitionsRequestHistoriesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RequestHistoriesClient.WorkflowRunActionRepetitionsRequestHistoriesList`

```go
ctx := context.TODO()
id := requesthistories.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "workflowName", "runName", "actionName", "repetitionName")

// alternatively `client.WorkflowRunActionRepetitionsRequestHistoriesList(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsRequestHistoriesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
