
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechatmessagereply` Documentation

The `mechatmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechatmessagereply"
```


### Client Initialization

```go
client := mechatmessagereply.NewMeChatMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatMessageReplyClient.CreateMeChatByIdMessageByIdReply`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

payload := mechatmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateMeChatByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.CreateMeChatByIdMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeChatByIdMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.CreateMeChatByIdMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeChatByIdMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.DeleteMeChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteMeChatByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.GetMeChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetMeChatByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.GetMeChatByIdMessageByIdReplyCount`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.GetMeChatByIdMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.ListMeChatByIdMessageByIdReplies`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

// alternatively `client.ListMeChatByIdMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListMeChatByIdMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatMessageReplyClient.SetMeChatByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mechatmessagereply.SetMeChatByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetMeChatByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.UnsetMeChatByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mechatmessagereply.UnsetMeChatByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeChatByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageReplyClient.UpdateMeChatByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mechatmessagereply.NewMeChatMessageReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mechatmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateMeChatByIdMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
