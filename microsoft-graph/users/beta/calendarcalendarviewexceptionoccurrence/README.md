
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


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarCalendarViewExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.CreateCalendarCalendarViewExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrence.CreateCalendarCalendarViewExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.GetCalendarCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.GetCalendarCalendarViewExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceClient.ListCalendarCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrence.NewUserIdCalendarIdCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarCalendarViewExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
