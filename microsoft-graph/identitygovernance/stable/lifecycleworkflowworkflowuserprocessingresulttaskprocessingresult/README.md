
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId")

payload := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultCreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowId", "userProcessingResultId")

read, err := client.GetLifecycleWorkflowUserProcessingResultTaskCount(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowId", "userProcessingResultId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowId", "userProcessingResultId")

// alternatively `client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
