
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/triggers` Documentation

The `triggers` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/triggers"
```


### Client Initialization

```go
client := triggers.NewTriggersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

payload := triggers.Trigger{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.Delete`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.Get`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggersClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := triggers.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id, triggers.DefaultListByDataBoxEdgeDeviceOperationOptions())` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id, triggers.DefaultListByDataBoxEdgeDeviceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
