
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowworkflowtemplatetasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtemplatetasktaskprocessingresult` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/lifecycleworkflowworkflowtemplatetasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

payload := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResumeRequest{
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


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.GetLifecycleWorkflowTemplateTaskProcessingResult`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateIdValue", "taskIdValue", "taskProcessingResultIdValue")

read, err := client.GetLifecycleWorkflowTemplateTaskProcessingResult(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultGetLifecycleWorkflowTemplateTaskProcessingResultOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.GetLifecycleWorkflowTemplateTaskProcessingResultsCount`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

read, err := client.GetLifecycleWorkflowTemplateTaskProcessingResultsCount(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultGetLifecycleWorkflowTemplateTaskProcessingResultsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.ListLifecycleWorkflowTemplateTaskProcessingResults`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateIdValue", "taskIdValue")

// alternatively `client.ListLifecycleWorkflowTemplateTaskProcessingResults(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTemplateTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
