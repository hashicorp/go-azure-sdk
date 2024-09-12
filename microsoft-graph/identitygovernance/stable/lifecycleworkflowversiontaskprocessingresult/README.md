
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowversiontaskprocessingresult` Documentation

The `lifecycleworkflowversiontaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowversiontaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowversiontaskprocessingresult.NewLifecycleWorkflowVersionTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowVersionTaskProcessingResultClient.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowversiontaskprocessingresult.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowVersionTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowVersionTaskProcessingResultClient.GetLifecycleWorkflowVersionTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowVersionTaskProcessingResult(ctx, id, lifecycleworkflowversiontaskprocessingresult.DefaultGetLifecycleWorkflowVersionTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowVersionTaskProcessingResultClient.GetLifecycleWorkflowVersionTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowVersionTaskProcessingResultsCount(ctx, id, lifecycleworkflowversiontaskprocessingresult.DefaultGetLifecycleWorkflowVersionTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowVersionTaskProcessingResultClient.ListLifecycleWorkflowVersionTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowversiontaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowVersionTaskProcessingResults(ctx, id, lifecycleworkflowversiontaskprocessingresult.DefaultListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowVersionTaskProcessingResultsComplete(ctx, id, lifecycleworkflowversiontaskprocessingresult.DefaultListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
