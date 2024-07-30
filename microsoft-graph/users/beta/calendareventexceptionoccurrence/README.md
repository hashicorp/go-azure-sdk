
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrence` Documentation

The `calendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrence"
```


### Client Initialization

```go
client := calendareventexceptionoccurrence.NewCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarEventExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.CreateCalendarEventExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrence.CreateCalendarEventExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.GetCalendarEventExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.GetCalendarEventExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarEventExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceClient.ListCalendarEventExceptionOccurrences`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrence.NewUserIdCalendarEventID("userIdValue", "eventIdValue")

// alternatively `client.ListCalendarEventExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
