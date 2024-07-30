
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

read, err := client.DeleteChatMessage(ctx, id)
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

read, err := client.GetChatMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.GetChatMessageCount`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatID("chatIdValue")

read, err := client.GetChatMessageCount(ctx, id)
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

// alternatively `client.ListChatMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListChatMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMessageClient.SetMeChatMessageReaction`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessage.SetMeChatMessageReactionRequest{
	// ...
}


read, err := client.SetMeChatMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMessageClient.UnsetMeChatMessageReaction`

```go
ctx := context.TODO()
id := chatmessage.NewMeChatIdMessageID("chatIdValue", "chatMessageIdValue")

payload := chatmessage.UnsetMeChatMessageReactionRequest{
	// ...
}


read, err := client.UnsetMeChatMessageReaction(ctx, id, payload)
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
