
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrence` Documentation

The `calendarcalendarviewexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrence.NewCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.AcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.AcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrence(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultAcceptCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CancelCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.CancelCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrence(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultCancelCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.DeclineCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.DeclineCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrence(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultDeclineCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.DismissCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.DismissCalendarViewExceptionOccurrenceReminder(ctx, id, calendarcalendarviewexceptionoccurrence.DefaultDismissCalendarViewExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.ForwardCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.ForwardCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrence(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultForwardCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrence(ctx, id, calendarcalendarviewexceptionoccurrence.DefaultGetCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

read, err := client.GetCalendarViewExceptionOccurrencesCount(ctx, id, calendarcalendarviewexceptionoccurrence.DefaultGetCalendarViewExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.ListCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

// alternatively `client.ListCalendarViewExceptionOccurrences(ctx, id, calendarcalendarviewexceptionoccurrence.DefaultListCalendarViewExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrencesComplete(ctx, id, calendarcalendarviewexceptionoccurrence.DefaultListCalendarViewExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.SnoozeCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.SnoozeCalendarViewExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceReminder(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultSnoozeCalendarViewExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.TentativelyAcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrence.TentativelyAcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrence(ctx, id, payload, calendarcalendarviewexceptionoccurrence.DefaultTentativelyAcceptCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
