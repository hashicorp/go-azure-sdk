
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowruntaskprocessingresult` Documentation

The `lifecycleworkflowruntaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowruntaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowruntaskprocessingresult.NewLifecycleWorkflowRunTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowRunTaskProcessingResultClient.CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowruntaskprocessingresult.CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowRunTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowRunTaskProcessingResult(ctx, id, lifecycleworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowRunTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowRunTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

read, err := client.GetLifecycleWorkflowRunTaskProcessingResultsCount(ctx, id, lifecycleworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowRunTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowRunTaskProcessingResultClient.ListLifecycleWorkflowRunTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

// alternatively `client.ListLifecycleWorkflowRunTaskProcessingResults(ctx, id, lifecycleworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowRunTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowRunTaskProcessingResultsComplete(ctx, id, lifecycleworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowRunTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
