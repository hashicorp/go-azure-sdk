
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowrunuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowrunuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowrunuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowRunUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowRunUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowRunUserProcessingResultTaskCount(ctx, id, lifecycleworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowRunUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowRunUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowRunUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowRunUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
