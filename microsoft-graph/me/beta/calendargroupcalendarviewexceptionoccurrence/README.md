
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewexceptionoccurrence` Documentation

The `calendargroupcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarviewexceptionoccurrence.NewCalendarGroupCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.AcceptCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.AcceptCalendarGroupCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.CancelCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.CancelCalendarGroupCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.DeclineCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.DeclineCalendarGroupCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.DismissCalendarGroupCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarViewExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.ForwardCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.ForwardCalendarGroupCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrence(ctx, id, calendargroupcalendarviewexceptionoccurrence.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarViewExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrencesCount(ctx, id, calendargroupcalendarviewexceptionoccurrence.DefaultGetCalendarGroupCalendarViewExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.ListCalendarGroupCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrences(ctx, id, calendargroupcalendarviewexceptionoccurrence.DefaultListCalendarGroupCalendarViewExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrencesComplete(ctx, id, calendargroupcalendarviewexceptionoccurrence.DefaultListCalendarGroupCalendarViewExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.SnoozeCalendarGroupCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.SnoozeCalendarGroupCalendarViewExceptionOccurrenceReminderRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceRequest{
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
