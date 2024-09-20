
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggerruns` Documentation

The `triggerruns` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggerruns"
```


### Client Initialization

```go
client := triggerruns.NewTriggerrunsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggerrunsClient.Cancel`

```go
ctx := context.TODO()
id := triggerruns.NewTriggerRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "triggerName", "runId")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggerrunsClient.QueryByFactory`

```go
ctx := context.TODO()
id := triggerruns.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName")

payload := triggerruns.RunFilterParameters{
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


### Example Usage: `TriggerrunsClient.Rerun`

```go
ctx := context.TODO()
id := triggerruns.NewTriggerRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "triggerName", "runId")

read, err := client.Rerun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
