
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggers` Documentation

The `triggers` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggers"
```


### Client Initialization

```go
client := triggers.NewTriggersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

payload := triggers.TriggerResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, triggers.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggersClient.Delete`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggersClient.GetEventSubscriptionStatus`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

read, err := client.GetEventSubscriptionStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggersClient.ListByFactory`

```go
ctx := context.TODO()
id := triggers.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

// alternatively `client.ListByFactory(ctx, id)` can be used to do batched pagination
items, err := client.ListByFactoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TriggersClient.QueryByFactory`

```go
ctx := context.TODO()
id := triggers.NewFactoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue")

payload := triggers.TriggerFilterParameters{
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


### Example Usage: `TriggersClient.Start`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

if err := client.StartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.Stop`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

if err := client.StopThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.SubscribeToEvents`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

if err := client.SubscribeToEventsThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.UnsubscribeFromEvents`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

if err := client.UnsubscribeFromEventsThenPoll(ctx, id); err != nil {
	// handle the error
}
```
