
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflow` Documentation

The `lifecycleworkflowdeleteditemworkflow` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflow"
```


### Client Initialization

```go
client := lifecycleworkflowdeleteditemworkflow.NewLifecycleWorkflowDeletedItemWorkflowClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivate`

```go
ctx := context.TODO()
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowIdValue")

payload := lifecycleworkflowdeleteditemworkflow.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivateRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceActivate(ctx, id, payload)
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
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowIdValue")

payload := lifecycleworkflowdeleteditemworkflow.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersion(ctx, id, payload)
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
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowIdValue")

read, err := client.CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestore(ctx, id)
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
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowIdValue")

read, err := client.DeleteLifecycleWorkflowDeletedItemWorkflow(ctx, id)
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
id := lifecycleworkflowdeleteditemworkflow.NewIdentityGovernanceLifecycleWorkflowDeletedItemWorkflowID("workflowIdValue")

read, err := client.GetLifecycleWorkflowDeletedItemWorkflow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowDeletedItemWorkflowClient.GetLifecycleWorkflowDeletedItemWorkflowCount`

```go
ctx := context.TODO()


read, err := client.GetLifecycleWorkflowDeletedItemWorkflowCount(ctx)
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


// alternatively `client.ListLifecycleWorkflowDeletedItemWorkflows(ctx)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowDeletedItemWorkflowsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
