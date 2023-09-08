
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrence` Documentation

The `usereventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrence"
```


### Client Initialization

```go
client := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.CreateUserByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrence.CreateUserByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.GetUserByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.GetUserByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceClient.ListUserByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usereventexceptionoccurrence.NewUserEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
