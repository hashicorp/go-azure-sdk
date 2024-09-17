
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/logicapps` Documentation

The `logicapps` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/logicapps"
```


### Client Initialization

```go
client := logicapps.NewLogicAppsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogicAppsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

payload := logicapps.LogicApp{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.Delete`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.DeployWorkflowArtifacts`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

payload := logicapps.WorkflowArtifacts{
	// ...
}


read, err := client.DeployWorkflowArtifacts(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.Get`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.GetWorkflow`

```go
ctx := context.TODO()
id := logicapps.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue", "workflowValue")

read, err := client.GetWorkflow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.Invoke`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

read, err := client.Invoke(ctx, id, logicapps.DefaultInvokeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogicAppsClient.ListWorkflows`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

// alternatively `client.ListWorkflows(ctx, id)` can be used to do batched pagination
items, err := client.ListWorkflowsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LogicAppsClient.ListWorkflowsConnections`

```go
ctx := context.TODO()
id := logicapps.NewLogicAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "logicAppValue")

read, err := client.ListWorkflowsConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
