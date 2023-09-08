
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceinstance` Documentation

The `groupeventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstance.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceClient.ListGroupByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
