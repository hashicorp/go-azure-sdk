
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultID("workflowIdValue", "runIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
