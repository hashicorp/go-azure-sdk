
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chat` Documentation

The `chat` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/chat"
```


### Client Initialization

```go
client := chat.NewChatClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatClient.CreateChat`

```go
ctx := context.TODO()

payload := chat.Chat{
	// ...
}


read, err := client.CreateChat(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.DeleteChat`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

read, err := client.DeleteChat(ctx, id, chat.DefaultDeleteChatOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.GetChat`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

read, err := client.GetChat(ctx, id, chat.DefaultGetChatOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.GetChatsCount`

```go
ctx := context.TODO()


read, err := client.GetChatsCount(ctx, chat.DefaultGetChatsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.HideChatForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.HideChatForUserRequest{
	// ...
}


read, err := client.HideChatForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.ListChats`

```go
ctx := context.TODO()


// alternatively `client.ListChats(ctx, chat.DefaultListChatsOperationOptions())` can be used to do batched pagination
items, err := client.ListChatsComplete(ctx, chat.DefaultListChatsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatClient.MarkChatReadForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.MarkChatReadForUserRequest{
	// ...
}


read, err := client.MarkChatReadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.MarkChatUnreadForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.MarkChatUnreadForUserRequest{
	// ...
}


read, err := client.MarkChatUnreadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.RemoveChatAllAccessForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.RemoveChatAllAccessForUserRequest{
	// ...
}


read, err := client.RemoveChatAllAccessForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.SendChatActivityNotification`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.SendChatActivityNotificationRequest{
	// ...
}


read, err := client.SendChatActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.UnhideChatForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.UnhideChatForUserRequest{
	// ...
}


read, err := client.UnhideChatForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.UpdateChat`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.Chat{
	// ...
}


read, err := client.UpdateChat(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
