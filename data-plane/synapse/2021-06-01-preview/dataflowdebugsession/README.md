
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflowdebugsession` Documentation

The `dataflowdebugsession` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflowdebugsession"
```


### Client Initialization

```go
client := dataflowdebugsession.NewDataFlowDebugSessionClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataFlowDebugSessionClient.AddDataFlow`

```go
ctx := context.TODO()

payload := dataflowdebugsession.DataFlowDebugPackage{
	// ...
}


read, err := client.AddDataFlow(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataFlowDebugSessionClient.CreateDataFlowDebugSession`

```go
ctx := context.TODO()

payload := dataflowdebugsession.CreateDataFlowDebugSessionRequest{
	// ...
}


if err := client.CreateDataFlowDebugSessionThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowDebugSessionClient.DeleteDataFlowDebugSession`

```go
ctx := context.TODO()

payload := dataflowdebugsession.DeleteDataFlowDebugSessionRequest{
	// ...
}


read, err := client.DeleteDataFlowDebugSession(ctx, payload)
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

payload := dataflowdebugsession.DataFlowDebugCommandRequest{
	// ...
}


if err := client.ExecuteCommandThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowDebugSessionClient.QueryDataFlowDebugSessionsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.QueryDataFlowDebugSessionsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.QueryDataFlowDebugSessionsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
