
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowruntaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdTaskProcessingResultID("workflowId", "runId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdTaskProcessingResultID("workflowId", "runId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowId", "runId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowId", "runId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
