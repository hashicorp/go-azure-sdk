
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstance` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
