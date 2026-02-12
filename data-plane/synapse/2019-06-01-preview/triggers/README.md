
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/triggers` Documentation

The `triggers` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/triggers"
```


### Client Initialization

```go
client := triggers.NewTriggersClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggersClient.TriggerCreateOrUpdateTrigger`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

payload := triggers.TriggerResource{
	// ...
}


if err := client.TriggerCreateOrUpdateTriggerThenPoll(ctx, id, payload, triggers.DefaultTriggerCreateOrUpdateTriggerOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.TriggerDeleteTrigger`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

if err := client.TriggerDeleteTriggerThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.TriggerGetEventSubscriptionStatus`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

read, err := client.TriggerGetEventSubscriptionStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggersClient.TriggerGetTriggersByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.TriggerGetTriggersByWorkspace(ctx)` can be used to do batched pagination
items, err := client.TriggerGetTriggersByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TriggersClient.TriggerStartTrigger`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

if err := client.TriggerStartTriggerThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.TriggerStopTrigger`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

if err := client.TriggerStopTriggerThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.TriggerSubscribeTriggerToEvents`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

if err := client.TriggerSubscribeTriggerToEventsThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggersClient.TriggerUnsubscribeTriggerFromEvents`

```go
ctx := context.TODO()
id := triggers.NewTriggerID("triggerName")

if err := client.TriggerUnsubscribeTriggerFromEventsThenPoll(ctx, id); err != nil {
	// handle the error
}
```
