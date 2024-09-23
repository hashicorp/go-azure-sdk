
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmember` Documentation

The `joinedteamchannelmember` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmember"
```


### Client Initialization

```go
client := joinedteamchannelmember.NewJoinedTeamChannelMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMemberClient.AddJoinedTeamChannelMembers`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannelmember.AddJoinedTeamChannelMembersRequest{
	// ...
}


// alternatively `client.AddJoinedTeamChannelMembers(ctx, id, payload, joinedteamchannelmember.DefaultAddJoinedTeamChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddJoinedTeamChannelMembersComplete(ctx, id, payload, joinedteamchannelmember.DefaultAddJoinedTeamChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMemberClient.CreateJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamChannelMember(ctx, id, payload, joinedteamchannelmember.DefaultCreateJoinedTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.DeleteJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelIdMemberID("userId", "teamId", "channelId", "conversationMemberId")

read, err := client.DeleteJoinedTeamChannelMember(ctx, id, joinedteamchannelmember.DefaultDeleteJoinedTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.GetJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelIdMemberID("userId", "teamId", "channelId", "conversationMemberId")

read, err := client.GetJoinedTeamChannelMember(ctx, id, joinedteamchannelmember.DefaultGetJoinedTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.GetJoinedTeamChannelMembersCount`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.GetJoinedTeamChannelMembersCount(ctx, id, joinedteamchannelmember.DefaultGetJoinedTeamChannelMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.ListJoinedTeamChannelMembers`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

// alternatively `client.ListJoinedTeamChannelMembers(ctx, id, joinedteamchannelmember.DefaultListJoinedTeamChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMembersComplete(ctx, id, joinedteamchannelmember.DefaultListJoinedTeamChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMemberClient.UpdateJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewUserIdJoinedTeamIdChannelIdMemberID("userId", "teamId", "channelId", "conversationMemberId")

payload := joinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamChannelMember(ctx, id, payload, joinedteamchannelmember.DefaultUpdateJoinedTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
