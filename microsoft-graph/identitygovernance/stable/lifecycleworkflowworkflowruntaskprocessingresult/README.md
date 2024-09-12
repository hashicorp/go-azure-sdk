
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


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowruntaskprocessingresult.CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest{
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


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowRunTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultID("workflowIdValue", "runIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowRunTaskProcessingResult(ctx, id, lifecycleworkflowworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowRunTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.GetLifecycleWorkflowRunTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

read, err := client.GetLifecycleWorkflowRunTaskProcessingResultsCount(ctx, id, lifecycleworkflowworkflowruntaskprocessingresult.DefaultGetLifecycleWorkflowRunTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowRunTaskProcessingResultClient.ListLifecycleWorkflowRunTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowruntaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdRunID("workflowIdValue", "runIdValue")

// alternatively `client.ListLifecycleWorkflowRunTaskProcessingResults(ctx, id, lifecycleworkflowworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowRunTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowRunTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowruntaskprocessingresult.DefaultListLifecycleWorkflowRunTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
