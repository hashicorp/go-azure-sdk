
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceinstance` Documentation

The `meeventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstance.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.GetMeEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.GetMeEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceClient.ListMeEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
