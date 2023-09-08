
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrence` Documentation

The `groupeventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrence"
```


### Client Initialization

```go
client := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrence.CreateGroupByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.GetGroupByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.GetGroupByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceClient.ListGroupByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrence.NewGroupEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
