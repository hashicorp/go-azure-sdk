
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelines` Documentation

The `pipelines` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelines"
```


### Client Initialization

```go
client := pipelines.NewPipelinesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `PipelinesClient.PipelineCreateOrUpdatePipeline`

```go
ctx := context.TODO()
id := pipelines.NewPipelineID("pipelineName")

payload := pipelines.PipelineResource{
	// ...
}


if err := client.PipelineCreateOrUpdatePipelineThenPoll(ctx, id, payload, pipelines.DefaultPipelineCreateOrUpdatePipelineOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `PipelinesClient.PipelineCreatePipelineRun`

```go
ctx := context.TODO()
id := pipelines.NewPipelineID("pipelineName")
var payload map[string]interface{}

read, err := client.PipelineCreatePipelineRun(ctx, id, payload, pipelines.DefaultPipelineCreatePipelineRunOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelinesClient.PipelineDeletePipeline`

```go
ctx := context.TODO()
id := pipelines.NewPipelineID("pipelineName")

if err := client.PipelineDeletePipelineThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PipelinesClient.PipelineGetPipeline`

```go
ctx := context.TODO()
id := pipelines.NewPipelineID("pipelineName")

read, err := client.PipelineGetPipeline(ctx, id, pipelines.DefaultPipelineGetPipelineOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelinesClient.PipelineGetPipelinesByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.PipelineGetPipelinesByWorkspace(ctx)` can be used to do batched pagination
items, err := client.PipelineGetPipelinesByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PipelinesClient.PipelineRenamePipeline`

```go
ctx := context.TODO()
id := pipelines.NewPipelineID("pipelineName")

payload := pipelines.ArtifactRenameRequest{
	// ...
}


if err := client.PipelineRenamePipelineThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
