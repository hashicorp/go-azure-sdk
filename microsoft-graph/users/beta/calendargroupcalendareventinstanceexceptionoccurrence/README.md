
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrence` Documentation

The `calendargroupcalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendareventinstanceexceptionoccurrence.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.ListCalendarGroupCalendarEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
