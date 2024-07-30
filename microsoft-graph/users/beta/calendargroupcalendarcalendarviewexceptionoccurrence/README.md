
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrence` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrence.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceClient.ListCalendarGroupCalendarCalendarViewExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarCalendarViewExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
