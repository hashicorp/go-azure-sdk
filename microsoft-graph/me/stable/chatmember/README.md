
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmember` Documentation

The `chatmember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/chatmember"
```


### Client Initialization

```go
client := chatmember.NewChatMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChatMemberClient.AddChatMembers`

```go
ctx := context.TODO()
id := chatmember.NewMeChatID("chatIdValue")

payload := chatmember.AddChatMembersRequest{
	// ...
}


// alternatively `client.AddChatMembers(ctx, id, payload, chatmember.DefaultAddChatMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddChatMembersComplete(ctx, id, payload, chatmember.DefaultAddChatMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ChatMemberClient.CreateChatMember`

```go
ctx := context.TODO()
id := chatmember.NewMeChatID("chatIdValue")

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
id := chatmember.NewMeChatIdMemberID("chatIdValue", "conversationMemberIdValue")

read, err := client.DeleteChatMember(ctx, id, chatmember.DefaultDeleteChatMemberOperationOptions())
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
id := chatmember.NewMeChatIdMemberID("chatIdValue", "conversationMemberIdValue")

read, err := client.GetChatMember(ctx, id, chatmember.DefaultGetChatMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ChatMemberClient.GetChatMembersCount`

```go
ctx := context.TODO()
id := chatmember.NewMeChatID("chatIdValue")

read, err := client.GetChatMembersCount(ctx, id, chatmember.DefaultGetChatMembersCountOperationOptions())
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
id := chatmember.NewMeChatID("chatIdValue")

// alternatively `client.ListChatMembers(ctx, id, chatmember.DefaultListChatMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListChatMembersComplete(ctx, id, chatmember.DefaultListChatMembersOperationOptions())
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
id := chatmember.NewMeChatIdMemberID("chatIdValue", "conversationMemberIdValue")

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
