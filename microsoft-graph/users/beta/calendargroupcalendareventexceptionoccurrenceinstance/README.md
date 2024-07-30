
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstance` Documentation

The `calendargroupcalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrenceinstance.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.ListCalendarGroupCalendarEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
