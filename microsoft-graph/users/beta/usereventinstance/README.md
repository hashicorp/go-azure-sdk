
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstance` Documentation

The `usereventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstance"
```


### Client Initialization

```go
client := usereventinstance.NewUserEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.CreateUserByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstance.CreateUserByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.GetUserByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.GetUserByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceClient.ListUserByIdEventByIdInstances`

```go
ctx := context.TODO()
id := usereventinstance.NewUserEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
