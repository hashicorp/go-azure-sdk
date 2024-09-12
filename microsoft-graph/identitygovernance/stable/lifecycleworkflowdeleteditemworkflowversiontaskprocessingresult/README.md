
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCount(ctx, id, lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowversiontaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
