
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
