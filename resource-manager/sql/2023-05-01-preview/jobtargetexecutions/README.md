
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/jobtargetexecutions` Documentation

The `jobtargetexecutions` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/jobtargetexecutions"
```


### Client Initialization

```go
client := jobtargetexecutions.NewJobTargetExecutionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobTargetExecutionsClient.Get`

```go
ctx := context.TODO()
id := jobtargetexecutions.NewTargetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "jobAgentValue", "jobValue", "jobExecutionIdValue", "stepValue", "targetIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobTargetExecutionsClient.ListByJobExecution`

```go
ctx := context.TODO()
id := jobtargetexecutions.NewExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "jobAgentValue", "jobValue", "jobExecutionIdValue")

// alternatively `client.ListByJobExecution(ctx, id, jobtargetexecutions.DefaultListByJobExecutionOperationOptions())` can be used to do batched pagination
items, err := client.ListByJobExecutionComplete(ctx, id, jobtargetexecutions.DefaultListByJobExecutionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JobTargetExecutionsClient.ListByStep`

```go
ctx := context.TODO()
id := jobtargetexecutions.NewExecutionStepID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "jobAgentValue", "jobValue", "jobExecutionIdValue", "stepValue")

// alternatively `client.ListByStep(ctx, id, jobtargetexecutions.DefaultListByStepOperationOptions())` can be used to do batched pagination
items, err := client.ListByStepComplete(ctx, id, jobtargetexecutions.DefaultListByStepOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
