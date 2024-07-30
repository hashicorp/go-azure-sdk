
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


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrence.CreateCalendarGroupCalendarEventExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceTentativelyAccept(ctx, id, payload)
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

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceClient.GetCalendarGroupCalendarEventExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceCount(ctx, id)
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

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
