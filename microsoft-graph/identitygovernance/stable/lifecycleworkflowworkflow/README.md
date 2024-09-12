
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


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflow`

```go
ctx := context.TODO()

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.CreateLifecycleWorkflow(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.CreateLifecycleWorkflowIdentityGovernanceActivate`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowIdentityGovernanceActivateRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowIdentityGovernanceActivate(ctx, id, payload)
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
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.CreateLifecycleWorkflowIdentityGovernanceCreateNewVersionRequest{
	// ...
}


read, err := client.CreateLifecycleWorkflowIdentityGovernanceCreateNewVersion(ctx, id, payload)
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
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.CreateLifecycleWorkflowIdentityGovernanceRestore(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.DeleteLifecycleWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.DeleteLifecycleWorkflow(ctx, id, lifecycleworkflowworkflow.DefaultDeleteLifecycleWorkflowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.GetLifecycleWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

read, err := client.GetLifecycleWorkflow(ctx, id, lifecycleworkflowworkflow.DefaultGetLifecycleWorkflowOperationOptions())
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


### Example Usage: `LifecycleWorkflowWorkflowClient.ListLifecycleWorkflows`

```go
ctx := context.TODO()


// alternatively `client.ListLifecycleWorkflows(ctx, lifecycleworkflowworkflow.DefaultListLifecycleWorkflowsOperationOptions())` can be used to do batched pagination
items, err := client.ListLifecycleWorkflowsComplete(ctx, lifecycleworkflowworkflow.DefaultListLifecycleWorkflowsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LifecycleWorkflowWorkflowClient.UpdateLifecycleWorkflow`

```go
ctx := context.TODO()
id := lifecycleworkflowworkflow.NewIdentityGovernanceLifecycleWorkflowWorkflowID("workflowIdValue")

payload := lifecycleworkflowworkflow.IdentityGovernanceWorkflow{
	// ...
}


read, err := client.UpdateLifecycleWorkflow(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
