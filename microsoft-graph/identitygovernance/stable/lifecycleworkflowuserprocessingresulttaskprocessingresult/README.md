
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowuserprocessingresulttaskprocessingresult` Documentation

The `lifecycleworkflowuserprocessingresulttaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowuserprocessingresulttaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowUserProcessingResultTaskCount(ctx, id, lifecycleworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowUserProcessingResultTaskCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowUserProcessingResultTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowUserProcessingResultTaskProcessingResult(ctx, id, lifecycleworkflowuserprocessingresulttaskprocessingresult.DefaultGetLifecycleWorkflowUserProcessingResultTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowUserProcessingResultTaskProcessingResultClient.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
