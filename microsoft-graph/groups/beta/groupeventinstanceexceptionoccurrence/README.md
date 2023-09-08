
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventinstanceexceptionoccurrence` Documentation

The `groupeventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrence.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceClient.ListGroupByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
