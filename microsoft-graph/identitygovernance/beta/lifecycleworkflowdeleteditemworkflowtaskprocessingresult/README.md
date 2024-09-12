
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowtaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowtaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowtaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowtaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowtaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
