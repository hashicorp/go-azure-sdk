
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmessagereply` Documentation

The `chatmessagereply` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmessagereply"
```


### Client Initialization

```go
client := chatmessagereply.NewChatMessageReplyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatMessageReplyClient.CreateChatMessageReply`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageID("chatId", "chatMessageId")

payload := chatmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateChatMessageReply(ctx, id, payload, chatmessagereply.DefaultCreateChatMessageReplyOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

read, err := client.CreateChatMessageReplySoftDelete(ctx, id, chatmessagereply.DefaultCreateChatMessageReplySoftDeleteOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

read, err := client.CreateChatMessageReplyUndoSoftDelete(ctx, id, chatmessagereply.DefaultCreateChatMessageReplyUndoSoftDeleteOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

read, err := client.DeleteChatMessageReply(ctx, id, chatmessagereply.DefaultDeleteChatMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.GetChatMessageRepliesCount`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageID("chatId", "chatMessageId")

read, err := client.GetChatMessageRepliesCount(ctx, id, chatmessagereply.DefaultGetChatMessageRepliesCountOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

read, err := client.GetChatMessageReply(ctx, id, chatmessagereply.DefaultGetChatMessageReplyOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageID("chatId", "chatMessageId")

// alternatively `client.ListChatMessageReplies(ctx, id, chatmessagereply.DefaultListChatMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListChatMessageRepliesComplete(ctx, id, chatmessagereply.DefaultListChatMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMessageReplyClient.SetChatMessageReplyReaction`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

payload := chatmessagereply.SetChatMessageReplyReactionRequest{
	// ...
}


read, err := client.SetChatMessageReplyReaction(ctx, id, payload, chatmessagereply.DefaultSetChatMessageReplyReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageReplyClient.UnsetChatMessageReplyReaction`

```go
ctx := context.TODO()
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

payload := chatmessagereply.UnsetChatMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetChatMessageReplyReaction(ctx, id, payload, chatmessagereply.DefaultUnsetChatMessageReplyReactionOperationOptions())
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
id := chatmessagereply.NewMeChatIdMessageIdReplyID("chatId", "chatMessageId", "chatMessageId1")

payload := chatmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateChatMessageReply(ctx, id, payload, chatmessagereply.DefaultUpdateChatMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
