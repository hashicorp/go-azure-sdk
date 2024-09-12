
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstance` Documentation

The `eventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := eventexceptionoccurrenceinstance.NewEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.AcceptEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.AcceptEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CancelEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CancelEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.DeclineEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.DeclineEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.DismissEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissEventExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.ForwardEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.ForwardEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.GetEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetEventExceptionOccurrenceInstance(ctx, id, eventexceptionoccurrenceinstance.DefaultGetEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.GetEventExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventExceptionOccurrenceInstancesCount(ctx, id, eventexceptionoccurrenceinstance.DefaultGetEventExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.ListEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListEventExceptionOccurrenceInstances(ctx, id, eventexceptionoccurrenceinstance.DefaultListEventExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceInstancesComplete(ctx, id, eventexceptionoccurrenceinstance.DefaultListEventExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.SnoozeEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.SnoozeEventExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeEventExceptionOccurrenceInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.TentativelyAcceptEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.TentativelyAcceptEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
