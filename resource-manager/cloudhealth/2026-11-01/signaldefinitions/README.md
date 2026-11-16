
## `github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/signaldefinitions` Documentation

The `signaldefinitions` SDK allows for interaction with Azure Resource Manager `cloudhealth` (API Version `2026-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/signaldefinitions"
```


### Client Initialization

```go
client := signaldefinitions.NewSignalDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SignalDefinitionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := signaldefinitions.NewSignaldefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "signaldefinitionName")

payload := signaldefinitions.SignalDefinition{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SignalDefinitionsClient.Delete`

```go
ctx := context.TODO()
id := signaldefinitions.NewSignaldefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "signaldefinitionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SignalDefinitionsClient.Get`

```go
ctx := context.TODO()
id := signaldefinitions.NewSignaldefinitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "signaldefinitionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignalDefinitionsClient.ListByHealthModel`

```go
ctx := context.TODO()
id := signaldefinitions.NewHealthmodelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName")

// alternatively `client.ListByHealthModel(ctx, id, signaldefinitions.DefaultListByHealthModelOperationOptions())` can be used to do batched pagination
items, err := client.ListByHealthModelComplete(ctx, id, signaldefinitions.DefaultListByHealthModelOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
