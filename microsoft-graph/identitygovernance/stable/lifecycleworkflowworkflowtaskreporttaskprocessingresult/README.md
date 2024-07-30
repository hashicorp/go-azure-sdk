
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtaskreporttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowtaskreporttaskprocessingresult.CreateLifecycleWorkflowWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowWorkflowTaskReportTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultID("workflowIdValue", "taskReportIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowTaskReportTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.GetLifecycleWorkflowWorkflowTaskReportTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

read, err := client.GetLifecycleWorkflowWorkflowTaskReportTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportID("workflowIdValue", "taskReportIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowTaskReportTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
