
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmessage` Documentation

The `userchatmessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmessage"
```


### Client Initialization

```go
client := userchatmessage.NewUserChatMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatMessageClient.CreateUserByIdChatByIdMessage`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatmessage.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdChatByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.CreateUserByIdChatByIdMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdChatByIdMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.CreateUserByIdChatByIdMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdChatByIdMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.DeleteUserByIdChatByIdMessageById`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

read, err := client.DeleteUserByIdChatByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.GetUserByIdChatByIdMessageById`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdChatByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.GetUserByIdChatByIdMessageCount`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.GetUserByIdChatByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.ListUserByIdChatByIdMessages`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.ListUserByIdChatByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatMessageClient.SetUserByIdChatByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

payload := userchatmessage.SetUserByIdChatByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdChatByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.UnsetUserByIdChatByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

payload := userchatmessage.UnsetUserByIdChatByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdChatByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageClient.UpdateUserByIdChatByIdMessageById`

```go
ctx := context.TODO()
id := userchatmessage.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

payload := userchatmessage.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdChatByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
