
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrence` Documentation

The `calendarviewinstanceexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendarviewinstanceexceptionoccurrence.NewCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.AcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.AcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultAcceptCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CancelCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.CancelCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultCancelCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.DeclineCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.DeclineCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultDeclineCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.DismissCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultDismissCalendarViewInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.ForwardCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.ForwardCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultForwardCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewInstanceExceptionOccurrence(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewInstanceExceptionOccurrencesCount(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.ListCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewInstanceExceptionOccurrences(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrencesComplete(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.SnoozeCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.SnoozeCalendarViewInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultSnoozeCalendarViewInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrence.TentativelyAcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarviewinstanceexceptionoccurrence.DefaultTentativelyAcceptCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
