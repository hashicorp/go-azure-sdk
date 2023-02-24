
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/pipelineruns` Documentation

The `pipelineruns` SDK allows for interaction with the Azure Resource Manager Service `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/pipelineruns"
```


### Client Initialization

```go
client := pipelineruns.NewPipelineRunsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PipelineRunsClient.Cancel`

```go
ctx := context.TODO()
id := pipelineruns.NewPipelineRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "runIdValue")

read, err := client.Cancel(ctx, id, pipelineruns.DefaultCancelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelineRunsClient.Get`

```go
ctx := context.TODO()
id := pipelineruns.NewPipelineRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "runIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PipelineRunsClient.QueryByFactory`

```go
ctx := context.TODO()
id := pipelineruns.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := pipelineruns.RunFilterParameters{
	// ...
}


read, err := client.QueryByFactory(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
