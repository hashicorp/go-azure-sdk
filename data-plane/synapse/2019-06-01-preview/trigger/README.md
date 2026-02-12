
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/trigger` Documentation

The `trigger` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/trigger"
```


### Client Initialization

```go
client := trigger.NewTriggerClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggerClient.GetTrigger`

```go
ctx := context.TODO()
id := trigger.NewTriggerID("triggerName")

read, err := client.GetTrigger(ctx, id, trigger.DefaultGetTriggerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
