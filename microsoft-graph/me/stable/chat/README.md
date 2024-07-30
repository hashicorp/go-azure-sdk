
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chat` Documentation

The `chat` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chat"
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


### Example Usage: `ChatClient.CreateChatHideForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.CreateChatHideForUserRequest{
	// ...
}


read, err := client.CreateChatHideForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.CreateChatMarkChatReadForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.CreateChatMarkChatReadForUserRequest{
	// ...
}


read, err := client.CreateChatMarkChatReadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.CreateChatMarkChatUnreadForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.CreateChatMarkChatUnreadForUserRequest{
	// ...
}


read, err := client.CreateChatMarkChatUnreadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.CreateChatSendActivityNotification`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.CreateChatSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateChatSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.CreateChatUnhideForUser`

```go
ctx := context.TODO()
id := chat.NewMeChatID("chatIdValue")

payload := chat.CreateChatUnhideForUserRequest{
	// ...
}


read, err := client.CreateChatUnhideForUser(ctx, id, payload)
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

read, err := client.DeleteChat(ctx, id)
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

read, err := client.GetChat(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatClient.GetChatCount`

```go
ctx := context.TODO()


read, err := client.GetChatCount(ctx)
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


// alternatively `client.ListChats(ctx)` can be used to do batched pagination
items, err := client.ListChatsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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
