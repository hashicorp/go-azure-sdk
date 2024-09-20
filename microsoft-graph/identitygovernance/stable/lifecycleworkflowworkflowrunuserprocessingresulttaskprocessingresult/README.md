
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowId", "runId", "userProcessingResultId", "taskProcessingResultId")

payload := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.DefaultCreateLifecycleWorkflowRunUserProcessingResultTaskIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowRunUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId")

read, err := client.GetLifecycleWorkflowRunUserProcessingResultTaskCount(ctx, id, lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowRunUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowRunUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowId", "runId", "userProcessingResultId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowRunUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowRunUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowId", "runId", "userProcessingResultId")

// alternatively `client.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowRunUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
