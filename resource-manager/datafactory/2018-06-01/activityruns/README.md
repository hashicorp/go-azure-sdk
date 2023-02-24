
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/activityruns` Documentation

The `activityruns` SDK allows for interaction with the Azure Resource Manager Service `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/activityruns"
```


### Client Initialization

```go
client := activityruns.NewActivityrunsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ActivityrunsClient.QueryByPipelineRun`

```go
ctx := context.TODO()
id := activityruns.NewPipelineRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "runIdValue")

payload := activityruns.RunFilterParameters{
	// ...
}


read, err := client.QueryByPipelineRun(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
