
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowId", "taskReportId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowId", "taskReportId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
