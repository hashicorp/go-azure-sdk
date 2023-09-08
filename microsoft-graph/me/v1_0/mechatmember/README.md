
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatmember` Documentation

The `mechatmember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatmember"
```


### Client Initialization

```go
client := mechatmember.NewMeChatMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatMemberClient.AddMeChatByIdMember`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatID("chatIdValue")

payload := mechatmember.AddMeChatByIdMemberRequest{
	// ...
}


read, err := client.AddMeChatByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMemberClient.CreateMeChatByIdMember`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatID("chatIdValue")

payload := mechatmember.ConversationMember{
	// ...
}


read, err := client.CreateMeChatByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMemberClient.DeleteMeChatByIdMemberById`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatMemberID("chatIdValue", "conversationMemberIdValue")

read, err := client.DeleteMeChatByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMemberClient.GetMeChatByIdMemberById`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatMemberID("chatIdValue", "conversationMemberIdValue")

read, err := client.GetMeChatByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMemberClient.GetMeChatByIdMemberCount`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatID("chatIdValue")

read, err := client.GetMeChatByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatMemberClient.ListMeChatByIdMembers`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatID("chatIdValue")

// alternatively `client.ListMeChatByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListMeChatByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatMemberClient.UpdateMeChatByIdMemberById`

```go
ctx := context.TODO()
id := mechatmember.NewMeChatMemberID("chatIdValue", "conversationMemberIdValue")

payload := mechatmember.ConversationMember{
	// ...
}


read, err := client.UpdateMeChatByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
