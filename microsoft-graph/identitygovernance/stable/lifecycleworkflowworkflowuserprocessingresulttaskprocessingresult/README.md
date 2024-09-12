
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


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.CreateLifecycleWorkflowUserProcessingResultTaskIdentityGovernanceResumeRequest{
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


### Example Usage: `LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.GetLifecycleWorkflowUserProcessingResultTaskCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

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
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultIdTaskProcessingResultID("workflowIdValue", "userProcessingResultIdValue", "taskProcessingResultIdValue")

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
id := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowIdUserProcessingResultID("workflowIdValue", "userProcessingResultIdValue")

// alternatively `client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResults(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowUserProcessingResultTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.DefaultListLifecycleWorkflowUserProcessingResultTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
