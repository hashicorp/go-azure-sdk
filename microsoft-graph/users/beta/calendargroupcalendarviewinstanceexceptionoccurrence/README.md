
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewinstanceexceptionoccurrence` Documentation

The `calendargroupcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarviewinstanceexceptionoccurrence.NewCalendarGroupCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.AcceptCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.AcceptCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.CancelCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.CancelCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.DeclineCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.DeclineCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.DismissCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.ForwardCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.ForwardCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrence(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrencesCount(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.ListCalendarGroupCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceExceptionOccurrences(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceExceptionOccurrencesComplete(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.SnoozeCalendarGroupCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.SnoozeCalendarGroupCalendarViewInstanceExceptionOccurrenceReminderRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarViewInstanceExceptionOccurrenceRequest{
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
