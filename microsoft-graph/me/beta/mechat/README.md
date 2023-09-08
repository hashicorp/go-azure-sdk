
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechat` Documentation

The `mechat` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mechat"
```


### Client Initialization

```go
client := mechat.NewMeChatClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatClient.CreateMeChat`

```go
ctx := context.TODO()

payload := mechat.Chat{
	// ...
}


read, err := client.CreateMeChat(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.CreateMeChatByIdHideForUser`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.CreateMeChatByIdHideForUserRequest{
	// ...
}


read, err := client.CreateMeChatByIdHideForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.CreateMeChatByIdMarkChatReadForUser`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.CreateMeChatByIdMarkChatReadForUserRequest{
	// ...
}


read, err := client.CreateMeChatByIdMarkChatReadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.CreateMeChatByIdMarkChatUnreadForUser`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.CreateMeChatByIdMarkChatUnreadForUserRequest{
	// ...
}


read, err := client.CreateMeChatByIdMarkChatUnreadForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.CreateMeChatByIdSendActivityNotification`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.CreateMeChatByIdSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateMeChatByIdSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.CreateMeChatByIdUnhideForUser`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.CreateMeChatByIdUnhideForUserRequest{
	// ...
}


read, err := client.CreateMeChatByIdUnhideForUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.DeleteMeChatById`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

read, err := client.DeleteMeChatById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.GetMeChatById`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

read, err := client.GetMeChatById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.GetMeChatCount`

```go
ctx := context.TODO()


read, err := client.GetMeChatCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatClient.ListMeChats`

```go
ctx := context.TODO()


// alternatively `client.ListMeChats(ctx)` can be used to do batched pagination
items, err := client.ListMeChatsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatClient.UpdateMeChatById`

```go
ctx := context.TODO()
id := mechat.NewMeChatID("chatIdValue")

payload := mechat.Chat{
	// ...
}


read, err := client.UpdateMeChatById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
