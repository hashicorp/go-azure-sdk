
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowId", "runId", "userProcessingResultId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskCount(ctx, id, lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowId", "runId", "userProcessingResultId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
