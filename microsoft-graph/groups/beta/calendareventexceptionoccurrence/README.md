
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrence` Documentation

The `calendareventexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrence"
```


### Client Initialization

```go
client := calendareventexceptionoccurrence.NewCalendarEventExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.AcceptCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.AcceptCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarEventExceptionOccurrence(ctx, id, payload, calendareventexceptionoccurrence.DefaultAcceptCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CancelCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.CancelCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarEventExceptionOccurrence(ctx, id, payload, calendareventexceptionoccurrence.DefaultCancelCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.DeclineCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.DeclineCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarEventExceptionOccurrence(ctx, id, payload, calendareventexceptionoccurrence.DefaultDeclineCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.DismissCalendarEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.DismissCalendarEventExceptionOccurrenceReminder(ctx, id, calendareventexceptionoccurrence.DefaultDismissCalendarEventExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.ForwardCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.ForwardCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarEventExceptionOccurrence(ctx, id, payload, calendareventexceptionoccurrence.DefaultForwardCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.GetCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarEventExceptionOccurrence(ctx, id, calendareventexceptionoccurrence.DefaultGetCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.GetCalendarEventExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.GetCalendarEventExceptionOccurrencesCount(ctx, id, calendareventexceptionoccurrence.DefaultGetCalendarEventExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.ListCalendarEventExceptionOccurrences`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventID("groupId", "eventId")

// alternatively `client.ListCalendarEventExceptionOccurrences(ctx, id, calendareventexceptionoccurrence.DefaultListCalendarEventExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrencesComplete(ctx, id, calendareventexceptionoccurrence.DefaultListCalendarEventExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.SnoozeCalendarEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.SnoozeCalendarEventExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventExceptionOccurrenceReminder(ctx, id, payload, calendareventexceptionoccurrence.DefaultSnoozeCalendarEventExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.TentativelyAcceptCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendareventexceptionoccurrence.TentativelyAcceptCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventExceptionOccurrence(ctx, id, payload, calendareventexceptionoccurrence.DefaultTentativelyAcceptCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
