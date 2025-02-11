
## `github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/experimentexecutions` Documentation

The `experimentexecutions` SDK allows for interaction with Azure Resource Manager `chaosstudio` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/chaosstudio/2025-01-01/experimentexecutions"
```


### Client Initialization

```go
client := experimentexecutions.NewExperimentExecutionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExperimentExecutionsClient.ExperimentsExecutionDetails`

```go
ctx := context.TODO()
id := experimentexecutions.NewExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentName", "executionId")

read, err := client.ExperimentsExecutionDetails(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentExecutionsClient.ExperimentsGetExecution`

```go
ctx := context.TODO()
id := experimentexecutions.NewExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentName", "executionId")

read, err := client.ExperimentsGetExecution(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExperimentExecutionsClient.ExperimentsListAllExecutions`

```go
ctx := context.TODO()
id := experimentexecutions.NewExperimentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "experimentName")

// alternatively `client.ExperimentsListAllExecutions(ctx, id)` can be used to do batched pagination
items, err := client.ExperimentsListAllExecutionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
