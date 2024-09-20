
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrence` Documentation

The `calendarviewexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrence.NewCalendarViewExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.AcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.AcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrence(ctx, id, payload, calendarviewexceptionoccurrence.DefaultAcceptCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CancelCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.CancelCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrence(ctx, id, payload, calendarviewexceptionoccurrence.DefaultCancelCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.DeclineCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.DeclineCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrence(ctx, id, payload, calendarviewexceptionoccurrence.DefaultDeclineCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.DismissCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.DismissCalendarViewExceptionOccurrenceReminder(ctx, id, calendarviewexceptionoccurrence.DefaultDismissCalendarViewExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.ForwardCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.ForwardCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrence(ctx, id, payload, calendarviewexceptionoccurrence.DefaultForwardCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrence(ctx, id, calendarviewexceptionoccurrence.DefaultGetCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewID("groupId", "eventId")

read, err := client.GetCalendarViewExceptionOccurrencesCount(ctx, id, calendarviewexceptionoccurrence.DefaultGetCalendarViewExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.ListCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewID("groupId", "eventId")

// alternatively `client.ListCalendarViewExceptionOccurrences(ctx, id, calendarviewexceptionoccurrence.DefaultListCalendarViewExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrencesComplete(ctx, id, calendarviewexceptionoccurrence.DefaultListCalendarViewExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.SnoozeCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.SnoozeCalendarViewExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceReminder(ctx, id, payload, calendarviewexceptionoccurrence.DefaultSnoozeCalendarViewExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.TentativelyAcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrence.TentativelyAcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrence(ctx, id, payload, calendarviewexceptionoccurrence.DefaultTentativelyAcceptCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
