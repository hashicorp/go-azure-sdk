
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowruntaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowruntaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowruntaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
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
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

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
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowIdValue", "runIdValue")

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
id := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunID("workflowIdValue", "runIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
