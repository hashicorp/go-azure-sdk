
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload)
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
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

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
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

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
id := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
