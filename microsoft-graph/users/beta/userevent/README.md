
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userevent` Documentation

The `userevent` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userevent"
```


### Client Initialization

```go
client := userevent.NewUserEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventClient.CreateUserByIdEvent`

```go
ctx := context.TODO()
id := userevent.NewUserID("userIdValue")

payload := userevent.Event{
	// ...
}


read, err := client.CreateUserByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdAccept`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdCancel`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdDecline`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.CreateUserByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdForward`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.CreateUserByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.CreateUserByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.DeleteUserByIdEventById`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.DeleteUserByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.GetUserByIdEventById`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.GetUserByIdEventCount`

```go
ctx := context.TODO()
id := userevent.NewUserID("userIdValue")

read, err := client.GetUserByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventClient.ListUserByIdEvents`

```go
ctx := context.TODO()
id := userevent.NewUserID("userIdValue")

// alternatively `client.ListUserByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserEventClient.UpdateUserByIdEventById`

```go
ctx := context.TODO()
id := userevent.NewUserEventID("userIdValue", "eventIdValue")

payload := userevent.Event{
	// ...
}


read, err := client.UpdateUserByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
