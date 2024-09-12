
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chat` Documentation

The `chat` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chat"
```


### Client Initialization

```go
client := chat.NewChatClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatClient.CreateChat`

```go
ctx := context.TODO()
id := chat.NewUserID("userIdValue")

payload := chat.Chat{
	// ...
}


read, err := client.CreateChat(ctx, id, payload)
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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserID("userIdValue")

read, err := client.GetChatsCount(ctx, id, chat.DefaultGetChatsCountOperationOptions())
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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserID("userIdValue")

// alternatively `client.ListChats(ctx, id, chat.DefaultListChatsOperationOptions())` can be used to do batched pagination
items, err := client.ListChatsComplete(ctx, id, chat.DefaultListChatsOperationOptions())
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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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


### Example Usage: `ChatClient.SendChatActivityNotification`

```go
ctx := context.TODO()
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
