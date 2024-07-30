
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
