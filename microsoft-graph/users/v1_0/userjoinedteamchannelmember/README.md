
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmember` Documentation

The `userjoinedteamchannelmember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmember"
```


### Client Initialization

```go
client := userjoinedteamchannelmember.NewUserJoinedTeamChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamChannelMemberClient.AddUserByIdJoinedTeamByIdChannelByIdMember`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := userjoinedteamchannelmember.AddUserByIdJoinedTeamByIdChannelByIdMemberRequest{
	// ...
}


read, err := client.AddUserByIdJoinedTeamByIdChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.CreateUserByIdJoinedTeamByIdChannelByIdMember`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := userjoinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.DeleteUserByIdJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelMemberID("userIdValue", "teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.GetUserByIdJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelMemberID("userIdValue", "teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.GetUserByIdJoinedTeamByIdChannelByIdMemberCount`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.ListUserByIdJoinedTeamByIdChannelByIdMembers`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdChannelByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdChannelByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamChannelMemberClient.UpdateUserByIdJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmember.NewUserJoinedTeamChannelMemberID("userIdValue", "teamIdValue", "channelIdValue", "conversationMemberIdValue")

payload := userjoinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdChannelByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
