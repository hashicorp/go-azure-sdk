
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId")

payload := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultCreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultID("workflowId", "userProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskCount(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultID("workflowId", "userProcessingResultId")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
