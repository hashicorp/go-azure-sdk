
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowworkflowtaskreporttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtaskreporttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowworkflowtaskreporttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId")

payload := lifecycleworkflowworkflowtaskreporttaskprocessingresult.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowworkflowtaskreporttaskprocessingresult.DefaultCreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowTaskReportTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowId", "taskReportId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowTaskReportTaskProcessingResult(ctx, id, lifecycleworkflowworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowTaskReportTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowTaskReportTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowId", "taskReportId")

read, err := client.GetLifecycleWorkflowTaskReportTaskProcessingResultsCount(ctx, id, lifecycleworkflowworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowTaskReportTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.ListLifecycleWorkflowTaskReportTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowId", "taskReportId")

// alternatively `client.ListLifecycleWorkflowTaskReportTaskProcessingResults(ctx, id, lifecycleworkflowworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowTaskReportTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTaskReportTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowTaskReportTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
