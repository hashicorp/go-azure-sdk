
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrence` Documentation

The `eventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrence"
```


### Client Initialization

```go
client := eventexceptionoccurrence.NewEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateEventExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CreateEventExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventexceptionoccurrence.CreateEventExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.GetEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.GetEventExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.GetEventExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.ListEventExceptionOccurrences`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewGroupIdEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListEventExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
