
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmessagereply` Documentation

The `chatmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmessagereply"
```


### Client Initialization

```go
client := chatmessagereply.NewChatMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatMessageReplyClient.CreateChatMessageReply`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateChatMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.CreateChatMessageReplySoftDelete`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateChatMessageReplySoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.CreateChatMessageReplyUndoSoftDelete`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateChatMessageReplyUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.DeleteChatMessageReply`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteChatMessageReply(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.GetChatMessageReply`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetChatMessageReply(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.GetChatMessageReplyCount`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.GetChatMessageReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.ListChatMessageReplies`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

// alternatively `client.ListChatMessageReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListChatMessageRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMessageReplyClient.SetMeChatMessageReplyReaction`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := chatmessagereply.SetMeChatMessageReplyReactionRequest{
	// ...
}


read, err := client.SetMeChatMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.UnsetMeChatMessageReplyReaction`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := chatmessagereply.UnsetMeChatMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetMeChatMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.UpdateChatMessageReply`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := chatmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateChatMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
