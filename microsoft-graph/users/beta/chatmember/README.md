
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmember` Documentation

The `chatmember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmember"
```


### Client Initialization

```go
client := chatmember.NewChatMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatMemberClient.AddUserChatMember`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatID("userIdValue", "chatIdValue")

payload := chatmember.AddUserChatMemberRequest{
	// ...
}


read, err := client.AddUserChatMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.CreateChatMember`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatID("userIdValue", "chatIdValue")

payload := chatmember.ConversationMember{
	// ...
}


read, err := client.CreateChatMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.DeleteChatMember`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatIdMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

read, err := client.DeleteChatMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.GetChatMember`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatIdMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

read, err := client.GetChatMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.GetChatMemberCount`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatID("userIdValue", "chatIdValue")

read, err := client.GetChatMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.ListChatMembers`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatID("userIdValue", "chatIdValue")

// alternatively `client.ListChatMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListChatMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMemberClient.UpdateChatMember`

```go
ctx := context.TODO()
id := chatmember.NewUserIdChatIdMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

payload := chatmember.ConversationMember{
	// ...
}


read, err := client.UpdateChatMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
