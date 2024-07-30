
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
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
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultCount(ctx, id)
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
id := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
