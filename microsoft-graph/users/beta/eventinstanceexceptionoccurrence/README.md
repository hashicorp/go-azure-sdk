
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrence` Documentation

The `eventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := eventinstanceexceptionoccurrence.NewEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateEventInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CreateEventInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CreateEventInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.GetEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetEventInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.GetEventInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.ListEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListEventInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListEventInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
