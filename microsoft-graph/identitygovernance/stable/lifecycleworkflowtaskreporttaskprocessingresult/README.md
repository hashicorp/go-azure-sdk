
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowtaskreporttaskprocessingresult` Documentation

The `lifecycleworkflowtaskreporttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowtaskreporttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowTaskReportTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowTaskReportTaskProcessingResultClient.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowtaskreporttaskprocessingresult.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowTaskReportTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowTaskReportTaskProcessingResult(ctx, id, lifecycleworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowTaskReportTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowTaskReportTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

read, err := client.GetLifecycleWorkflowTaskReportTaskProcessingResultsCount(ctx, id, lifecycleworkflowtaskreporttaskprocessingresult.DefaultGetLifecycleWorkflowTaskReportTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTaskReportTaskProcessingResultClient.ListLifecycleWorkflowTaskReportTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

// alternatively `client.ListLifecycleWorkflowTaskReportTaskProcessingResults(ctx, id, lifecycleworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowTaskReportTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTaskReportTaskProcessingResultsComplete(ctx, id, lifecycleworkflowtaskreporttaskprocessingresult.DefaultListLifecycleWorkflowTaskReportTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
