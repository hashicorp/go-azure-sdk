
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelineruns` Documentation

The `pipelineruns` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelineruns"
```


### Client Initialization

```go
client := pipelineruns.NewPipelineRunsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `PipelineRunsClient.PipelineRunCancelPipelineRun`

```go
ctx := context.TODO()
id := pipelineruns.NewPipelineRunID("runId")

read, err := client.PipelineRunCancelPipelineRun(ctx, id, pipelineruns.DefaultPipelineRunCancelPipelineRunOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelineRunsClient.PipelineRunGetPipelineRun`

```go
ctx := context.TODO()
id := pipelineruns.NewPipelineRunID("runId")

read, err := client.PipelineRunGetPipelineRun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelineRunsClient.PipelineRunQueryPipelineRunsByWorkspace`

```go
ctx := context.TODO()

payload := pipelineruns.RunFilterParameters{
	// ...
}


read, err := client.PipelineRunQueryPipelineRunsByWorkspace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
