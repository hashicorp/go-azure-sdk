
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrenceinstance` Documentation

The `usereventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventexceptionoccurrenceinstance.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.GetUserByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.GetUserByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceInstanceClient.ListUserByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
