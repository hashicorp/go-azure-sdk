
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/activityruns` Documentation

The `activityruns` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/activityruns"
```


### Client Initialization

```go
client := activityruns.NewActivityrunsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `ActivityrunsClient.PipelineRunQueryActivityRuns`

```go
ctx := context.TODO()
id := activityruns.NewPipelinePipelineRunID("pipelineName", "runId")

payload := activityruns.RunFilterParameters{
	// ...
}


read, err := client.PipelineRunQueryActivityRuns(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
