
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewinstanceexceptionoccurrence` Documentation

The `calendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceexceptionoccurrence.NewCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrence.CreateCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarCalendarViewInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.GetCalendarCalendarViewInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceClient.ListCalendarCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarCalendarViewInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
