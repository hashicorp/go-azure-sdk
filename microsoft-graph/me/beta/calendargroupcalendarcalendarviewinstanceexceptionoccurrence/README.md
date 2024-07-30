
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `calendargroupcalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
