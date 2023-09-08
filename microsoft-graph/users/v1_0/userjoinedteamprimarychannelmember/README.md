
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmember` Documentation

The `userjoinedteamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmember"
```


### Client Initialization

```go
client := userjoinedteamprimarychannelmember.NewUserJoinedTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.AddUserByIdJoinedTeamByIdPrimaryChannelMember`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteamprimarychannelmember.AddUserByIdJoinedTeamByIdPrimaryChannelMemberRequest{
	// ...
}


read, err := client.AddUserByIdJoinedTeamByIdPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMember`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.DeleteUserByIdJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.GetUserByIdJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.GetUserByIdJoinedTeamByIdPrimaryChannelMemberCount`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.ListUserByIdJoinedTeamByIdPrimaryChannelMembers`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdPrimaryChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdPrimaryChannelMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMemberClient.UpdateUserByIdJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmember.NewUserJoinedTeamPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

payload := userjoinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdPrimaryChannelMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
