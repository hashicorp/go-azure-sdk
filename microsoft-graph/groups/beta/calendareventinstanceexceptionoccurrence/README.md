
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrence` Documentation

The `calendareventinstanceexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendareventinstanceexceptionoccurrence.NewCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.AcceptCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.AcceptCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarEventInstanceExceptionOccurrence(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultAcceptCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CancelCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.CancelCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarEventInstanceExceptionOccurrence(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultCancelCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.DeclineCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.DeclineCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarEventInstanceExceptionOccurrence(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultDeclineCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.DismissCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissCalendarEventInstanceExceptionOccurrenceReminder(ctx, id, calendareventinstanceexceptionoccurrence.DefaultDismissCalendarEventInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.ForwardCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.ForwardCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarEventInstanceExceptionOccurrence(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultForwardCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.GetCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarEventInstanceExceptionOccurrence(ctx, id, calendareventinstanceexceptionoccurrence.DefaultGetCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.GetCalendarEventInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarEventInstanceExceptionOccurrencesCount(ctx, id, calendareventinstanceexceptionoccurrence.DefaultGetCalendarEventInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.ListCalendarEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrences(ctx, id, calendareventinstanceexceptionoccurrence.DefaultListCalendarEventInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrencesComplete(ctx, id, calendareventinstanceexceptionoccurrence.DefaultListCalendarEventInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.SnoozeCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.SnoozeCalendarEventInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventInstanceExceptionOccurrenceReminder(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultSnoozeCalendarEventInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrence.TentativelyAcceptCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventInstanceExceptionOccurrence(ctx, id, payload, calendareventinstanceexceptionoccurrence.DefaultTentativelyAcceptCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
