
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflows` Documentation

The `dataflows` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflows"
```


### Client Initialization

```go
client := dataflows.NewDataFlowsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataFlowsClient.DataFlowCreateOrUpdateDataFlow`

```go
ctx := context.TODO()
id := dataflows.NewDataflowID("dataflowName")

payload := dataflows.DataFlowResource{
	// ...
}


if err := client.DataFlowCreateOrUpdateDataFlowThenPoll(ctx, id, payload, dataflows.DefaultDataFlowCreateOrUpdateDataFlowOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowsClient.DataFlowDeleteDataFlow`

```go
ctx := context.TODO()
id := dataflows.NewDataflowID("dataflowName")

if err := client.DataFlowDeleteDataFlowThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataFlowsClient.DataFlowGetDataFlow`

```go
ctx := context.TODO()
id := dataflows.NewDataflowID("dataflowName")

read, err := client.DataFlowGetDataFlow(ctx, id, dataflows.DefaultDataFlowGetDataFlowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataFlowsClient.DataFlowGetDataFlowsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.DataFlowGetDataFlowsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.DataFlowGetDataFlowsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataFlowsClient.DataFlowRenameDataFlow`

```go
ctx := context.TODO()
id := dataflows.NewDataflowID("dataflowName")

payload := dataflows.ArtifactRenameRequest{
	// ...
}


if err := client.DataFlowRenameDataFlowThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
