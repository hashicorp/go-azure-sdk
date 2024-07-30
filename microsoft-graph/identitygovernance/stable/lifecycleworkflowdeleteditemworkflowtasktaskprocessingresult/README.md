
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultID("workflowIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskID("workflowIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
