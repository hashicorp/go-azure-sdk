
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventinstanceexceptionoccurrence` Documentation

The `eventinstanceexceptionoccurrence` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := eventinstanceexceptionoccurrence.NewEventInstanceExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.AcceptEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.AcceptEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptEventInstanceExceptionOccurrence(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultAcceptEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CancelEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.CancelEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelEventInstanceExceptionOccurrence(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultCancelEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.DeclineEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.DeclineEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineEventInstanceExceptionOccurrence(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultDeclineEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.DismissEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

read, err := client.DismissEventInstanceExceptionOccurrenceReminder(ctx, id, eventinstanceexceptionoccurrence.DefaultDismissEventInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.ForwardEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.ForwardEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardEventInstanceExceptionOccurrence(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultForwardEventInstanceExceptionOccurrenceOperationOptions())
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
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

read, err := client.GetEventInstanceExceptionOccurrence(ctx, id, eventinstanceexceptionoccurrence.DefaultGetEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.GetEventInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceID("eventId", "eventId1")

read, err := client.GetEventInstanceExceptionOccurrencesCount(ctx, id, eventinstanceexceptionoccurrence.DefaultGetEventInstanceExceptionOccurrencesCountOperationOptions())
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
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceID("eventId", "eventId1")

// alternatively `client.ListEventInstanceExceptionOccurrences(ctx, id, eventinstanceexceptionoccurrence.DefaultListEventInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventInstanceExceptionOccurrencesComplete(ctx, id, eventinstanceexceptionoccurrence.DefaultListEventInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.SnoozeEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.SnoozeEventInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeEventInstanceExceptionOccurrenceReminder(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultSnoozeEventInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.TentativelyAcceptEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewMeEventIdInstanceIdExceptionOccurrenceID("eventId", "eventId1", "eventId2")

payload := eventinstanceexceptionoccurrence.TentativelyAcceptEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventInstanceExceptionOccurrence(ctx, id, payload, eventinstanceexceptionoccurrence.DefaultTentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
