
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrence` Documentation

The `calendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendareventinstanceexceptionoccurrence.NewCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarEventInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CreateCalendarEventInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CreateCalendarEventInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.GetCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.GetCalendarEventInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.ListCalendarEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
