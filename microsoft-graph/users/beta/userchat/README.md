
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchat` Documentation

The `userchat` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchat"
```


### Client Initialization

```go
client := userchat.NewUserChatClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatClient.CreateUserByIdChat`

```go
ctx := context.TODO()
id := userchat.NewUserID("userIdValue")

payload := userchat.Chat{
	// ...
}


read, err := client.CreateUserByIdChat(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.CreateUserByIdChatByIdHideForUser`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.CreateUserByIdChatByIdHideForUserRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdHideForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.CreateUserByIdChatByIdMarkChatReadForUser`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.CreateUserByIdChatByIdMarkChatReadForUserRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdMarkChatReadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.CreateUserByIdChatByIdMarkChatUnreadForUser`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.CreateUserByIdChatByIdMarkChatUnreadForUserRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdMarkChatUnreadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.CreateUserByIdChatByIdSendActivityNotification`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.CreateUserByIdChatByIdSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.CreateUserByIdChatByIdUnhideForUser`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.CreateUserByIdChatByIdUnhideForUserRequest{
	// ...
}


read, err := client.CreateUserByIdChatByIdUnhideForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.DeleteUserByIdChatById`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.DeleteUserByIdChatById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.GetUserByIdChatById`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.GetUserByIdChatById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.GetUserByIdChatCount`

```go
ctx := context.TODO()
id := userchat.NewUserID("userIdValue")

read, err := client.GetUserByIdChatCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatClient.ListUserByIdChats`

```go
ctx := context.TODO()
id := userchat.NewUserID("userIdValue")

// alternatively `client.ListUserByIdChats(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatClient.UpdateUserByIdChatById`

```go
ctx := context.TODO()
id := userchat.NewUserChatID("userIdValue", "chatIdValue")

payload := userchat.Chat{
	// ...
}


read, err := client.UpdateUserByIdChatById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
