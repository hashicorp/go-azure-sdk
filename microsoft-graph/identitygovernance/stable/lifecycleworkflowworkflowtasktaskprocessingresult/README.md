
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtasktaskprocessingresult.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTaskTaskProcessingResultClient.CreateLifecycleWorkflowTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowtasktaskprocessingresult.CreateLifecycleWorkflowTaskProcessingResultIdentityGovernanceResumeRequest{
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


### Example Usage: `LifecycleWorkflowWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowTaskProcessingResult(ctx, id, lifecycleworkflowworkflowtasktaskprocessingresult.DefaultGetLifecycleWorkflowTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID("workflowIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowTaskProcessingResultsCount(ctx, id, lifecycleworkflowworkflowtasktaskprocessingresult.DefaultGetLifecycleWorkflowTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskTaskProcessingResultClient.ListLifecycleWorkflowTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskID("workflowIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowTaskProcessingResults(ctx, id, lifecycleworkflowworkflowtasktaskprocessingresult.DefaultListLifecycleWorkflowTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowtasktaskprocessingresult.DefaultListLifecycleWorkflowTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
