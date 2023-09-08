
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronlinemeeting` Documentation

The `useronlinemeeting` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronlinemeeting"
```


### Client Initialization

```go
client := useronlinemeeting.NewUserOnlineMeetingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnlineMeetingClient.CreateUserByIdOnlineMeeting`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserID("userIdValue")

payload := useronlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.CreateUserByIdOnlineMeeting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnlineMeetingClient.CreateUserByIdOnlineMeetingCreateOrGet`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserID("userIdValue")

payload := useronlinemeeting.CreateUserByIdOnlineMeetingCreateOrGetRequest{
	// ...
}


read, err := client.CreateUserByIdOnlineMeetingCreateOrGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnlineMeetingClient.DeleteUserByIdOnlineMeetingById`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

read, err := client.DeleteUserByIdOnlineMeetingById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnlineMeetingClient.GetUserByIdOnlineMeetingById`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

read, err := client.GetUserByIdOnlineMeetingById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnlineMeetingClient.GetUserByIdOnlineMeetingCount`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserID("userIdValue")

read, err := client.GetUserByIdOnlineMeetingCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnlineMeetingClient.ListUserByIdOnlineMeetings`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserID("userIdValue")

// alternatively `client.ListUserByIdOnlineMeetings(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnlineMeetingsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnlineMeetingClient.UpdateUserByIdOnlineMeetingById`

```go
ctx := context.TODO()
id := useronlinemeeting.NewUserOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

payload := useronlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.UpdateUserByIdOnlineMeetingById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
