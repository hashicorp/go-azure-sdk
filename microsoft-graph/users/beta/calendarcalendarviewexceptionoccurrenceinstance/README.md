
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstance` Documentation

The `calendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceinstance.NewCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CreateCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarCalendarViewExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.ListCalendarCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarCalendarViewExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
