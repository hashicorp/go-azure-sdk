
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrence` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrence.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.AcceptCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.AcceptCalendarGroupCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CancelCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CancelCalendarGroupCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.DeclineCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.DeclineCalendarGroupCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.DismissCalendarGroupCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarViewExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.ForwardCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.ForwardCalendarGroupCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrence(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrence.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarViewExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrencesCount(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrence.DefaultGetCalendarGroupCalendarViewExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.ListCalendarGroupCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrences(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrence.DefaultListCalendarGroupCalendarViewExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrencesComplete(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrence.DefaultListCalendarGroupCalendarViewExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.SnoozeCalendarGroupCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.SnoozeCalendarGroupCalendarViewExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarViewExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
