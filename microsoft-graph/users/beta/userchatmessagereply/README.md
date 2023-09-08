
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmessagereply` Documentation

The `userchatmessagereply` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmessagereply"
```


### Client Initialization

```go
client := userchatmessagereply.NewUserChatMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatMessageReplyClient.CreateUserByIdChatByIdMessageByIdReply`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

payload := userchatmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdChatByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.CreateUserByIdChatByIdMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdChatByIdMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.CreateUserByIdChatByIdMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdChatByIdMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.DeleteUserByIdChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteUserByIdChatByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.GetUserByIdChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetUserByIdChatByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.GetUserByIdChatByIdMessageByIdReplyCount`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdChatByIdMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.ListUserByIdChatByIdMessageByIdReplies`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageID("userIdValue", "chatIdValue", "chatMessageIdValue")

// alternatively `client.ListUserByIdChatByIdMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatByIdMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatMessageReplyClient.SetUserByIdChatByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userchatmessagereply.SetUserByIdChatByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdChatByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.UnsetUserByIdChatByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userchatmessagereply.UnsetUserByIdChatByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdChatByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMessageReplyClient.UpdateUserByIdChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userchatmessagereply.NewUserChatMessageReplyID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userchatmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdChatByIdMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
