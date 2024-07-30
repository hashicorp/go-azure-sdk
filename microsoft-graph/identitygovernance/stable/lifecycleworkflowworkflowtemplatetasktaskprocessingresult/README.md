
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtemplatetasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.CreateLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.CreateLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultIdentityGovernanceResume(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.GetLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResult(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.GetLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResults(ctx, id)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
