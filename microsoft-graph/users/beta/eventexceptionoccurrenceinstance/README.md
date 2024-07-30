
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


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateEventExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CreateEventExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstance.CreateEventExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
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

read, err := client.GetEventExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.GetEventExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewUserIdEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventExceptionOccurrenceInstanceCount(ctx, id)
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

// alternatively `client.ListEventExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
