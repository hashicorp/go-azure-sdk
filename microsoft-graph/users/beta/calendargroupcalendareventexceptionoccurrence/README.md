
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrence` Documentation

The `calendargroupcalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrence.NewCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.AcceptCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.AcceptCalendarGroupCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarEventExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CancelCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CancelCalendarGroupCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarEventExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.DeclineCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.DeclineCalendarGroupCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarEventExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.DismissCalendarGroupCalendarEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarEventExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.ForwardCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.ForwardCalendarGroupCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarEventExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.GetCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrence(ctx, id, calendargroupcalendareventexceptionoccurrence.DefaultGetCalendarGroupCalendarEventExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.GetCalendarGroupCalendarEventExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrencesCount(ctx, id, calendargroupcalendareventexceptionoccurrence.DefaultGetCalendarGroupCalendarEventExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.ListCalendarGroupCalendarEventExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrences(ctx, id, calendargroupcalendareventexceptionoccurrence.DefaultListCalendarGroupCalendarEventExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrencesComplete(ctx, id, calendargroupcalendareventexceptionoccurrence.DefaultListCalendarGroupCalendarEventExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.SnoozeCalendarGroupCalendarEventExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.SnoozeCalendarGroupCalendarEventExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarEventExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
