
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrence` Documentation

The `calendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrence.NewCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.AcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.AcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrence(ctx, id, payload)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CancelCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrence(ctx, id, payload)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.DeclineCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrence(ctx, id, payload)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.DismissCalendarViewExceptionOccurrenceReminder(ctx, id)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.ForwardCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrence(ctx, id, payload)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

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
id := calendarviewexceptionoccurrence.NewMeCalendarViewID("eventIdValue")

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
id := calendarviewexceptionoccurrence.NewMeCalendarViewID("eventIdValue")

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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.SnoozeCalendarViewExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceReminder(ctx, id, payload)
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
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.TentativelyAcceptCalendarViewExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
