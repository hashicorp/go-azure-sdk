
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrence` Documentation

The `calendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrence.NewCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.AcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.AcceptCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CancelCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CancelCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.DeclineCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.DeclineCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.DismissCalendarViewExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarViewExceptionOccurrenceReminder(ctx, id)
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
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.ForwardCalendarViewExceptionOccurrenceRequest{
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

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
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

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
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewID("userIdValue", "eventIdValue")

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
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.SnoozeCalendarViewExceptionOccurrenceReminderRequest{
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.TentativelyAcceptCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.TentativelyAcceptCalendarViewExceptionOccurrenceRequest{
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
