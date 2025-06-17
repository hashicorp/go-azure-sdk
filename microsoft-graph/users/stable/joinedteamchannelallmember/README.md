
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelallmember` Documentation

The `joinedteamchannelallmember` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelallmember"
```


### Client Initialization

```go
client := joinedteamchannelallmember.NewJoinedTeamChannelAllMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelAllMemberClient.AddJoinedTeamChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannelallmember.AddJoinedTeamChannelAllMembersRequest{
	// ...
}


// alternatively `client.AddJoinedTeamChannelAllMembers(ctx, id, payload, joinedteamchannelallmember.DefaultAddJoinedTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddJoinedTeamChannelAllMembersComplete(ctx, id, payload, joinedteamchannelallmember.DefaultAddJoinedTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.CreateJoinedTeamChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannelallmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamChannelAllMember(ctx, id, payload, joinedteamchannelallmember.DefaultCreateJoinedTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.DeleteJoinedTeamChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelIdAllMemberID("userId", "teamId", "channelId", "conversationMemberId")

read, err := client.DeleteJoinedTeamChannelAllMember(ctx, id, joinedteamchannelallmember.DefaultDeleteJoinedTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.GetJoinedTeamChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelIdAllMemberID("userId", "teamId", "channelId", "conversationMemberId")

read, err := client.GetJoinedTeamChannelAllMember(ctx, id, joinedteamchannelallmember.DefaultGetJoinedTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.GetJoinedTeamChannelAllMembersCount`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.GetJoinedTeamChannelAllMembersCount(ctx, id, joinedteamchannelallmember.DefaultGetJoinedTeamChannelAllMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.ListJoinedTeamChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

// alternatively `client.ListJoinedTeamChannelAllMembers(ctx, id, joinedteamchannelallmember.DefaultListJoinedTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelAllMembersComplete(ctx, id, joinedteamchannelallmember.DefaultListJoinedTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.RemoveJoinedTeamChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannelallmember.RemoveJoinedTeamChannelAllMembersRequest{
	// ...
}


// alternatively `client.RemoveJoinedTeamChannelAllMembers(ctx, id, payload, joinedteamchannelallmember.DefaultRemoveJoinedTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveJoinedTeamChannelAllMembersComplete(ctx, id, payload, joinedteamchannelallmember.DefaultRemoveJoinedTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelAllMemberClient.UpdateJoinedTeamChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamchannelallmember.NewUserIdJoinedTeamIdChannelIdAllMemberID("userId", "teamId", "channelId", "conversationMemberId")

payload := joinedteamchannelallmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamChannelAllMember(ctx, id, payload, joinedteamchannelallmember.DefaultUpdateJoinedTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
