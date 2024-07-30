
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstance` Documentation

The `calendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendareventexceptionoccurrenceinstance.NewCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarEventExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CreateCalendarEventExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CreateCalendarEventExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.GetCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarEventIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.GetCalendarEventExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventExceptionOccurrenceInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.ListCalendarEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewUserIdCalendarEventIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
