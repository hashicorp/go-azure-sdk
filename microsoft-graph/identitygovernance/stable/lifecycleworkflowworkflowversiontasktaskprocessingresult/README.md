
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowversiontasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowversiontasktaskprocessingresult.CreateLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowWorkflowVersionTaskTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskIdTaskProcessingResultID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowVersionTaskTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.GetLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.ListLifecycleWorkflowWorkflowVersionTaskTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskID("workflowIdValue", "workflowVersionVersionNumberValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowVersionTaskTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
