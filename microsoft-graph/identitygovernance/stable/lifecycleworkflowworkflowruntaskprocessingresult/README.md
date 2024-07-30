
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresult` Documentation

The `lifecycleworkflowworkflowruntaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowruntaskprocessingresult.NewLifecycleWorkflowWorkflowRunTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowRunTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowruntaskprocessingresult.CreateLifecycleWorkflowWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowRunTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowWorkflowRunTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowRunTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowWorkflowRunTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

read, err := client.GetLifecycleWorkflowWorkflowRunTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.ListLifecycleWorkflowWorkflowRunTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowRunTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowRunTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
