
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowId", "taskId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowId", "taskId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowId", "taskId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowId", "taskId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
