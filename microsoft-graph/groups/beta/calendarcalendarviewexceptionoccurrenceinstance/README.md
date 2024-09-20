
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstance` Documentation

The `calendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceinstance.NewCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.AcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultAcceptCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CancelCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.CancelCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultCancelCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.DeclineCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultDeclineCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.DismissCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultDismissCalendarViewExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.ForwardCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultForwardCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewExceptionOccurrenceInstance(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.ListCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstances(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.SnoozeCalendarViewExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultSnoozeCalendarViewExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarcalendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload, calendarcalendarviewexceptionoccurrenceinstance.DefaultTentativelyAcceptCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
