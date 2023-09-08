
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceexceptionoccurrence` Documentation

The `usereventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrence.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceClient.ListUserByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrence.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
