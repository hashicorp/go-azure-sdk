
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
