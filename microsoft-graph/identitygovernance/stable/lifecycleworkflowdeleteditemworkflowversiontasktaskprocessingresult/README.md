
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowId", "workflowVersionVersionNumber", "taskId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowId", "workflowVersionVersionNumber", "taskId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
