
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrence` Documentation

The `calendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceexceptionoccurrence.NewCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.AcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.AcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultAcceptCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CancelCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.CancelCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultCancelCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.DeclineCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.DeclineCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultDeclineCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.DismissCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, calendarcalendarviewinstanceexceptionoccurrence.DefaultDismissCalendarViewInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.ForwardCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.ForwardCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultForwardCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewInstanceExceptionOccurrence(ctx, id, calendarcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewInstanceExceptionOccurrencesCount(ctx, id, calendarcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.ListCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewInstanceExceptionOccurrences(ctx, id, calendarcalendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrencesComplete(ctx, id, calendarcalendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.SnoozeCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.SnoozeCalendarViewInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultSnoozeCalendarViewInstanceExceptionOccurrenceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewinstanceexceptionoccurrence.TentativelyAcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload, calendarcalendarviewinstanceexceptionoccurrence.DefaultTentativelyAcceptCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
