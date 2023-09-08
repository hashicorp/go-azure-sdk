
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstance` Documentation

The `groupcalendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceClient.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
