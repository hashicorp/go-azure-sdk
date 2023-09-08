
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatmessage` Documentation

The `mechatmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatmessage"
```


### Client Initialization

```go
client := mechatmessage.NewMeChatMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatMessageClient.CreateMeChatByIdMessage`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatID("chatIdValue")

payload := mechatmessage.ChatMessage{
	// ...
}


read, err := client.CreateMeChatByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.CreateMeChatByIdMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.CreateMeChatByIdMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.CreateMeChatByIdMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.CreateMeChatByIdMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.DeleteMeChatByIdMessageById`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.DeleteMeChatByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.GetMeChatByIdMessageById`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

read, err := client.GetMeChatByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.GetMeChatByIdMessageCount`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatID("chatIdValue")

read, err := client.GetMeChatByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.ListMeChatByIdMessages`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatID("chatIdValue")

// alternatively `client.ListMeChatByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeChatByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatMessageClient.SetMeChatByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

payload := mechatmessage.SetMeChatByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.SetMeChatByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.UnsetMeChatByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

payload := mechatmessage.UnsetMeChatByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeChatByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMessageClient.UpdateMeChatByIdMessageById`

```go
ctx := context.TODO()
id := mechatmessage.NewMeChatMessageID("chatIdValue", "chatMessageIdValue")

payload := mechatmessage.ChatMessage{
	// ...
}


read, err := client.UpdateMeChatByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
