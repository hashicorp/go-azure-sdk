
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowtemplatetaskprocessingresult` Documentation

The `lifecycleworkflowtemplatetaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowtemplatetaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowtemplatetaskprocessingresult.NewLifecycleWorkflowTemplateTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowTemplateTaskProcessingResultClient.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowtemplatetaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowtemplatetaskprocessingresult.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTemplateTaskProcessingResultClient.GetLifecycleWorkflowTemplateTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowtemplatetaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowTemplateTaskProcessingResult(ctx, id, lifecycleworkflowtemplatetaskprocessingresult.DefaultGetLifecycleWorkflowTemplateTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTemplateTaskProcessingResultClient.GetLifecycleWorkflowTemplateTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowtemplatetaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowTemplateTaskProcessingResultsCount(ctx, id, lifecycleworkflowtemplatetaskprocessingresult.DefaultGetLifecycleWorkflowTemplateTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowTemplateTaskProcessingResultClient.ListLifecycleWorkflowTemplateTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowtemplatetaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowTemplateTaskProcessingResults(ctx, id, lifecycleworkflowtemplatetaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTemplateTaskProcessingResultsComplete(ctx, id, lifecycleworkflowtemplatetaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
