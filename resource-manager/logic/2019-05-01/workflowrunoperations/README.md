
## `github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowrunoperations` Documentation

The `workflowrunoperations` SDK allows for interaction with Azure Resource Manager `logic` (API Version `2019-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/workflowrunoperations"
```


### Client Initialization

```go
client := workflowrunoperations.NewWorkflowRunOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunOperationsClient.Get`

```go
ctx := context.TODO()
id := workflowrunoperations.NewOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workflowValue", "runValue", "operationIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
