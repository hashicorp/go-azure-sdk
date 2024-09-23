
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflow` Documentation

The `lifecycleworkflowworkflow` SDK allows for interaction with Microsoft Graph `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflow"
```


### Client Initialization

```go
client := lifecycleworkflowworkflow.NewLifecycleWorkflowWorkflowClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowIdentityGovernanceActivate`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowIdentityGovernanceActivateRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowIdentityGovernanceActivate(ctx, id, payload, lifecycleworkflowworkflow.DefaultCreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowIdentityGovernanceCreateNewVersion`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowIdentityGovernanceCreateNewVersionRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowIdentityGovernanceCreateNewVersion(ctx, id, payload, lifecycleworkflowworkflow.DefaultCreateLifecycleWorkflowIdentityGovernanceCreateNewVersionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowIdentityGovernanceRestore`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

read, err := client.CreateLifecycleWorkflowIdentityGovernanceRestore(ctx, id, lifecycleworkflowworkflow.DefaultCreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowWorkflow`

```go
ctx := context.TODO()

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflow(ctx, payload, lifecycleworkflowworkflow.DefaultCreateLifecycleWorkflowWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.DeleteLifecycleWorkflowWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

read, err := client.DeleteLifecycleWorkflowWorkflow(ctx, id, lifecycleworkflowworkflow.DefaultDeleteLifecycleWorkflowWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.GetLifecycleWorkflowCount`

```go
ctx := context.TODO()


read, err := client.GetLifecycleWorkflowCount(ctx, lifecycleworkflowworkflow.DefaultGetLifecycleWorkflowCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.GetLifecycleWorkflowWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

read, err := client.GetLifecycleWorkflowWorkflow(ctx, id, lifecycleworkflowworkflow.DefaultGetLifecycleWorkflowWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.ListLifecycleWorkflowWorkflows`

```go
ctx := context.TODO()


// alternatively `client.ListLifecycleWorkflowWorkflows(ctx, lifecycleworkflowworkflow.DefaultListLifecycleWorkflowWorkflowsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowsComplete(ctx, lifecycleworkflowworkflow.DefaultListLifecycleWorkflowWorkflowsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.UpdateLifecycleWorkflowWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowId")

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.UpdateLifecycleWorkflowWorkflow(ctx, id, payload, lifecycleworkflowworkflow.DefaultUpdateLifecycleWorkflowWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
