
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrence` Documentation

The `eventexceptionoccurrence` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrence"
```


### Client Initialization

```go
client := eventexceptionoccurrence.NewEventExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceClient.AcceptEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.AcceptEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptEventExceptionOccurrence(ctx, id, payload, eventexceptionoccurrence.DefaultAcceptEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.CancelEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.CancelEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelEventExceptionOccurrence(ctx, id, payload, eventexceptionoccurrence.DefaultCancelEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.DeclineEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.DeclineEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineEventExceptionOccurrence(ctx, id, payload, eventexceptionoccurrence.DefaultDeclineEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.DismissEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

read, err := client.DismissEventExceptionOccurrenceReminder(ctx, id, eventexceptionoccurrence.DefaultDismissEventExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.ForwardEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.ForwardEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardEventExceptionOccurrence(ctx, id, payload, eventexceptionoccurrence.DefaultForwardEventExceptionOccurrenceOperationOptions())
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
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

read, err := client.GetEventExceptionOccurrence(ctx, id, eventexceptionoccurrence.DefaultGetEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.GetEventExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventID("userId", "eventId")

read, err := client.GetEventExceptionOccurrencesCount(ctx, id, eventexceptionoccurrence.DefaultGetEventExceptionOccurrencesCountOperationOptions())
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
id := eventexceptionoccurrence.NewUserIdEventID("userId", "eventId")

// alternatively `client.ListEventExceptionOccurrences(ctx, id, eventexceptionoccurrence.DefaultListEventExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrencesComplete(ctx, id, eventexceptionoccurrence.DefaultListEventExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventExceptionOccurrenceClient.SnoozeEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.SnoozeEventExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeEventExceptionOccurrenceReminder(ctx, id, payload, eventexceptionoccurrence.DefaultSnoozeEventExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceClient.TentativelyAcceptEventExceptionOccurrence`

```go
ctx := context.TODO()
id := eventexceptionoccurrence.NewUserIdEventIdExceptionOccurrenceID("userId", "eventId", "eventId1")

payload := eventexceptionoccurrence.TentativelyAcceptEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventExceptionOccurrence(ctx, id, payload, eventexceptionoccurrence.DefaultTentativelyAcceptEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
