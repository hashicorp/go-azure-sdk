
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflow` Documentation

The `lifecycleworkflowworkflow` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflow"
```


### Client Initialization

```go
client := lifecycleworkflowworkflow.NewLifecycleWorkflowWorkflowClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowWorkflow`

```go
ctx := context.TODO()

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflow(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowWorkflowIdentityGovernanceActivate`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowWorkflowIdentityGovernanceActivateRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowIdentityGovernanceActivate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowWorkflowIdentityGovernanceCreateNewVersion`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowWorkflowIdentityGovernanceCreateNewVersionRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowWorkflowIdentityGovernanceCreateNewVersion(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowWorkflowIdentityGovernanceRestore`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.CreateLifecycleWorkflowWorkflowIdentityGovernanceRestore(ctx, id)
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
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.DeleteLifecycleWorkflowWorkflow(ctx, id)
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
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.GetLifecycleWorkflowWorkflow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.GetLifecycleWorkflowWorkflowCount`

```go
ctx := context.TODO()


read, err := client.GetLifecycleWorkflowWorkflowCount(ctx)
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


// alternatively `client.ListLifecycleWorkflowWorkflows(ctx)` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowWorkflowsComplete(ctx)
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
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.UpdateLifecycleWorkflowWorkflow(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
