
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstance` Documentation

The `calendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceinstance.NewCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceAccept`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceCancel`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceDecline`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarViewExceptionOccurrenceInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceForward`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CreateCalendarViewExceptionOccurrenceInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CreateCalendarViewExceptionOccurrenceInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstanceCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.ListCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
