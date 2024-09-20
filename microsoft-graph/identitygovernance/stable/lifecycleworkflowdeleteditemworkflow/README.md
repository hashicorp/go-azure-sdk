
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflow` Documentation

The `lifecycleworkflowdeleteditemworkflow` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflow"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflow.NewLifecycleWorkflowDeletedItemWorkflowClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivate`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowId")

payload := lifecycleworkflowdeleteditemworkflow.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivateRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivate(ctx, id, payload, lifecycleworkflowdeleteditemworkflow.DefaultCreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersion`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowId")

payload := lifecycleworkflowdeleteditemworkflow.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersion(ctx, id, payload, lifecycleworkflowdeleteditemworkflow.DefaultCreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestore`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowId")

read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestore(ctx, id, lifecycleworkflowdeleteditemworkflow.DefaultCreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.DeleteLifecycleWorkflowDeletedItemWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowId")

read, err := client.DeleteLifecycleWorkflowDeletedItemWorkflow(ctx, id, lifecycleworkflowdeleteditemworkflow.DefaultDeleteLifecycleWorkflowDeletedItemWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.GetLifecycleWorkflowDeletedItemWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowId")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflow(ctx, id, lifecycleworkflowdeleteditemworkflow.DefaultGetLifecycleWorkflowDeletedItemWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.GetLifecycleWorkflowDeletedItemWorkflowsCount`

```go
ctx := context.TODO()


read, err := client.GetLifecycleWorkflowDeletedItemWorkflowsCount(ctx, lifecycleworkflowdeleteditemworkflow.DefaultGetLifecycleWorkflowDeletedItemWorkflowsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.ListLifecycleWorkflowDeletedItemWorkflows`

```go
ctx := context.TODO()


// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflows(ctx, lifecycleworkflowdeleteditemworkflow.DefaultListLifecycleWorkflowDeletedItemWorkflowsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowsComplete(ctx, lifecycleworkflowdeleteditemworkflow.DefaultListLifecycleWorkflowDeletedItemWorkflowsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
