
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowversiontasktaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowId", "workflowVersionVersionNumber", "taskId", "taskProcessingResultId")

payload := lifecycleworkflowworkflowversiontasktaskprocessingresult.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowworkflowversiontasktaskprocessingresult.DefaultCreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResumeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowVersionTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowId", "workflowVersionVersionNumber", "taskId", "taskProcessingResultId")

read, err := client.GetLifecycleWorkflowVersionTaskProcessingResult(ctx, id, lifecycleworkflowworkflowversiontasktaskprocessingresult.DefaultGetLifecycleWorkflowVersionTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowVersionTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId")

read, err := client.GetLifecycleWorkflowVersionTaskProcessingResultsCount(ctx, id, lifecycleworkflowworkflowversiontasktaskprocessingresult.DefaultGetLifecycleWorkflowVersionTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.ListLifecycleWorkflowVersionTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowId", "workflowVersionVersionNumber", "taskId")

// alternatively `client.ListLifecycleWorkflowVersionTaskProcessingResults(ctx, id, lifecycleworkflowworkflowversiontasktaskprocessingresult.DefaultListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowVersionTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowversiontasktaskprocessingresult.DefaultListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
