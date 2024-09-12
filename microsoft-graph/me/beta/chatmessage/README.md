
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chatmessage` Documentation

The `chatmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chatmessage"
```


### Client Initialization

```go
client := chatmessage.NewChatMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatMessageClient.CreateChatMessage`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatID("chatIdValue")

payload := chatmessage.ChatMessage{
	// ...
}


read, err := client.CreateChatMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.CreateChatMessageSoftDelete`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.CreateChatMessageSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.CreateChatMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.CreateChatMessageUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.DeleteChatMessage`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.DeleteChatMessage(ctx, id, chatmessage.DefaultDeleteChatMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.GetChatMessage`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.GetChatMessage(ctx, id, chatmessage.DefaultGetChatMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.GetChatMessagesCount`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatID("chatIdValue")

read, err := client.GetChatMessagesCount(ctx, id, chatmessage.DefaultGetChatMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.ListChatMessages`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatID("chatIdValue")

// alternatively `client.ListChatMessages(ctx, id, chatmessage.DefaultListChatMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListChatMessagesComplete(ctx, id, chatmessage.DefaultListChatMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMessageClient.SetChatMessageReaction`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessage.SetChatMessageReactionRequest{
	// ...
}


read, err := client.SetChatMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.UnsetChatMessageReaction`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessage.UnsetChatMessageReactionRequest{
	// ...
}


read, err := client.UnsetChatMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.UpdateChatMessage`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessage.ChatMessage{
	// ...
}


read, err := client.UpdateChatMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
