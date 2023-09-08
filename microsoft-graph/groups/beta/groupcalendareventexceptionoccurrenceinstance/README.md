
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrenceinstance` Documentation

The `groupcalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventexceptionoccurrenceinstance.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceInstanceClient.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
