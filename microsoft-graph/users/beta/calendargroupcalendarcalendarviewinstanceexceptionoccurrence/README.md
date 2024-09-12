
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `calendargroupcalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.AcceptCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.AcceptCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CancelCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CancelCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.DeclineCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.DeclineCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.DismissCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.ForwardCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.ForwardCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrencesCount(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.ListCalendarGroupCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceExceptionOccurrences(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceExceptionOccurrencesComplete(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.SnoozeCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.SnoozeCalendarGroupCalendarViewInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
