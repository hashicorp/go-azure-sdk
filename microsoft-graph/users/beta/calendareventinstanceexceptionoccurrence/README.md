
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


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.AcceptCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.AcceptCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.CancelCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.CancelCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.DeclineCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.DeclineCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.DismissCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarEventInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.ForwardCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.ForwardCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
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
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventInstanceExceptionOccurrence(ctx, id, calendareventinstanceexceptionoccurrence.DefaultGetCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.GetCalendarEventInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventInstanceExceptionOccurrencesCount(ctx, id, calendareventinstanceexceptionoccurrence.DefaultGetCalendarEventInstanceExceptionOccurrencesCountOperationOptions())
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
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarIdEventIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrences(ctx, id, calendareventinstanceexceptionoccurrence.DefaultListCalendarEventInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrencesComplete(ctx, id, calendareventinstanceexceptionoccurrence.DefaultListCalendarEventInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.SnoozeCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.SnoozeCalendarEventInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventInstanceExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrence.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrence.TentativelyAcceptCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
