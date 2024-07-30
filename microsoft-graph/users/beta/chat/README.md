
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chat` Documentation

The `chat` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chat"
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


### Example Usage: `ChatClient.CreateChatHideForUser`

```go
ctx := context.TODO()
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

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
id := chat.NewUserID("userIdValue")

read, err := client.GetChatCount(ctx, id)
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

// alternatively `client.ListChats(ctx, id)` can be used to do batched pagination
items, err := client.ListChatsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatClient.RemoveUserChatAllAccessForUser`

```go
ctx := context.TODO()
id := chat.NewUserIdChatID("userIdValue", "chatIdValue")

payload := chat.RemoveUserChatAllAccessForUserRequest{
	// ...
}


read, err := client.RemoveUserChatAllAccessForUser(ctx, id, payload)
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
