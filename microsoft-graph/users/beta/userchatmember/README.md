
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmember` Documentation

The `userchatmember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userchatmember"
```


### Client Initialization

```go
client := userchatmember.NewUserChatMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatMemberClient.AddUserByIdChatByIdMember`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatmember.AddUserByIdChatByIdMemberRequest{
	// ...
}


read, err := client.AddUserByIdChatByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMemberClient.CreateUserByIdChatByIdMember`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatmember.ConversationMember{
	// ...
}


read, err := client.CreateUserByIdChatByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMemberClient.DeleteUserByIdChatByIdMemberById`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

read, err := client.DeleteUserByIdChatByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMemberClient.GetUserByIdChatByIdMemberById`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

read, err := client.GetUserByIdChatByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMemberClient.GetUserByIdChatByIdMemberCount`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.GetUserByIdChatByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatMemberClient.ListUserByIdChatByIdMembers`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.ListUserByIdChatByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatMemberClient.UpdateUserByIdChatByIdMemberById`

```go
ctx := context.TODO()
id := userchatmember.NewUserChatMemberID("userIdValue", "chatIdValue", "conversationMemberIdValue")

payload := userchatmember.ConversationMember{
	// ...
}


read, err := client.UpdateUserByIdChatByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
