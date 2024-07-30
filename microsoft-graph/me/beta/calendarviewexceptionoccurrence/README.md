
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


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.CreateCalendarViewExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.CreateCalendarViewExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrence.CreateCalendarViewExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceTentativelyAccept(ctx, id, payload)
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

read, err := client.GetCalendarViewExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceClient.GetCalendarViewExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrence.NewMeCalendarViewID("eventIdValue")

read, err := client.GetCalendarViewExceptionOccurrenceCount(ctx, id)
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

// alternatively `client.ListCalendarViewExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
