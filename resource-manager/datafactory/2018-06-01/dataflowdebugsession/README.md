
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflowdebugsession` Documentation

The `dataflowdebugsession` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflowdebugsession"
```


### Client Initialization

```go
client := dataflowdebugsession.NewDataFlowDebugSessionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataFlowDebugSessionClient.AddDataFlow`

```go
ctx := context.TODO()
id := dataflowdebugsession.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := dataflowdebugsession.DataFlowDebugPackage{
	// ...
}


read, err := client.AddDataFlow(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataFlowDebugSessionClient.Create`

```go
ctx := context.TODO()
id := dataflowdebugsession.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := dataflowdebugsession.CreateDataFlowDebugSessionRequest{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowDebugSessionClient.Delete`

```go
ctx := context.TODO()
id := dataflowdebugsession.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := dataflowdebugsession.DeleteDataFlowDebugSessionRequest{
	// ...
}


read, err := client.Delete(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataFlowDebugSessionClient.ExecuteCommand`

```go
ctx := context.TODO()
id := dataflowdebugsession.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := dataflowdebugsession.DataFlowDebugCommandRequest{
	// ...
}


if err := client.ExecuteCommandThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowDebugSessionClient.QueryByFactory`

```go
ctx := context.TODO()
id := dataflowdebugsession.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

// alternatively `client.QueryByFactory(ctx, id)` can be used to do batched pagination
items, err := client.QueryByFactoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
