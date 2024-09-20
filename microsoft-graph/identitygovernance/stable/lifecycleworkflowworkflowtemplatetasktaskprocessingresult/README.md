
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresult` Documentation

The `lifecycleworkflowworkflowtemplatetasktaskprocessingresult` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresult"
```


### Client Initialization

```go
client := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResume`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateId", "taskId", "taskProcessingResultId")

payload := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResumeRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResume(ctx, id, payload, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultCreateLifecycleWorkflowTemplateTaskProcessingResultIdentityGovernanceResumeOperationOptions())
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
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskIdTaskProcessingResultID("workflowTemplateId", "taskId", "taskProcessingResultId")

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
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateId", "taskId")

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
id := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewIdentityGovernanceLifecycleWorkflowWorkflowTemplateIdTaskID("workflowTemplateId", "taskId")

// alternatively `client.ListLifecycleWorkflowTemplateTaskProcessingResults(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowTemplateTaskProcessingResultsComplete(ctx, id, lifecycleworkflowworkflowtemplatetasktaskprocessingresult.DefaultListLifecycleWorkflowTemplateTaskProcessingResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
