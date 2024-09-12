
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowtaskprocessingresult` Documentation

The `lifecycleworkflowtaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowtaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowtaskprocessingresult.NewLifecycleWorkflowTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowTaskProcessingResultClient.CreateLifecycleWorkflowTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowtaskprocessingresult.CreateLifecycleWorkflowTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskProcessingResultClient.GetLifecycleWorkflowTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowTaskProcessingResult(ctx, id, lifecycleworkflowtaskprocessingresult.DefaultGetLifecycleWorkflowTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskProcessingResultClient.GetLifecycleWorkflowTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID("workflowIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowTaskProcessingResultsCount(ctx, id, lifecycleworkflowtaskprocessingresult.DefaultGetLifecycleWorkflowTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskProcessingResultClient.ListLifecycleWorkflowTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID("workflowIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowTaskProcessingResults(ctx, id, lifecycleworkflowtaskprocessingresult.DefaultListLifecycleWorkflowTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTaskProcessingResultsComplete(ctx, id, lifecycleworkflowtaskprocessingresult.DefaultListLifecycleWorkflowTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
