
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstance` Documentation

The `calendarviewexceptionoccurrenceinstance` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceinstance.NewCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.AcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultAcceptCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CancelCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.CancelCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultCancelCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.DeclineCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultDeclineCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.DismissCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultDismissCalendarViewExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.ForwardCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultForwardCalendarViewExceptionOccurrenceInstanceOperationOptions())
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
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewExceptionOccurrenceInstance(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
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
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstances(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.SnoozeCalendarViewExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultSnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarviewexceptionoccurrenceinstance.DefaultTentativelyAcceptCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
