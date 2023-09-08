
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrenceinstance` Documentation

The `groupcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewexceptionoccurrenceinstance.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceInstanceClient.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
