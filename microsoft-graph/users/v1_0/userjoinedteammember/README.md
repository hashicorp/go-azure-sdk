
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteammember` Documentation

The `userjoinedteammember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteammember"
```


### Client Initialization

```go
client := userjoinedteammember.NewUserJoinedTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamMemberClient.AddUserByIdJoinedTeamByIdMember`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteammember.AddUserByIdJoinedTeamByIdMemberRequest{
	// ...
}


read, err := client.AddUserByIdJoinedTeamByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamMemberClient.CreateUserByIdJoinedTeamByIdMember`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteammember.ConversationMember{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamMemberClient.DeleteUserByIdJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamMemberClient.GetUserByIdJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.GetUserByIdJoinedTeamByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamMemberClient.GetUserByIdJoinedTeamByIdMemberCount`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamMemberClient.ListUserByIdJoinedTeamByIdMembers`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamMemberClient.UpdateUserByIdJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteammember.NewUserJoinedTeamMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

payload := userjoinedteammember.ConversationMember{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
