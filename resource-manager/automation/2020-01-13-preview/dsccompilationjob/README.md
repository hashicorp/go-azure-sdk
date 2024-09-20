
## `github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/dsccompilationjob` Documentation

The `dsccompilationjob` SDK allows for interaction with Azure Resource Manager `automation` (API Version `2020-01-13-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/dsccompilationjob"
```


### Client Initialization

```go
client := dsccompilationjob.NewDscCompilationJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DscCompilationJobClient.Create`

```go
ctx := context.TODO()
id := commonids.NewAutomationCompilationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName", "compilationJobId")

payload := dsccompilationjob.DscCompilationJobCreateParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DscCompilationJobClient.Get`

```go
ctx := context.TODO()
id := commonids.NewAutomationCompilationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName", "compilationJobId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DscCompilationJobClient.GetStream`

```go
ctx := context.TODO()
id := dsccompilationjob.NewStreamID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName", "jobId", "jobStreamId")

read, err := client.GetStream(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DscCompilationJobClient.ListByAutomationAccount`

```go
ctx := context.TODO()
id := dsccompilationjob.NewAutomationAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName")

// alternatively `client.ListByAutomationAccount(ctx, id, dsccompilationjob.DefaultListByAutomationAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByAutomationAccountComplete(ctx, id, dsccompilationjob.DefaultListByAutomationAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DscCompilationJobClient.StreamListByJob`

```go
ctx := context.TODO()
id := commonids.NewAutomationCompilationJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountName", "compilationJobId")

// alternatively `client.StreamListByJob(ctx, id)` can be used to do batched pagination
items, err := client.StreamListByJobComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
