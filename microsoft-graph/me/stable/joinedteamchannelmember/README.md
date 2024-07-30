
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmember` Documentation

The `joinedteamchannelmember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmember"
```


### Client Initialization

```go
client := joinedteamchannelmember.NewJoinedTeamChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMemberClient.AddMeJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

payload := joinedteamchannelmember.AddMeJoinedTeamChannelMemberRequest{
	// ...
}


read, err := client.AddMeJoinedTeamChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.CreateJoinedTeamChannelMember`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

payload := joinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamChannelMember(ctx, id, payload)
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
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelIdMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.DeleteJoinedTeamChannelMember(ctx, id)
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
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelIdMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.GetJoinedTeamChannelMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMemberClient.GetJoinedTeamChannelMemberCount`

```go
ctx := context.TODO()
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

read, err := client.GetJoinedTeamChannelMemberCount(ctx, id)
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
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

// alternatively `client.ListJoinedTeamChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMembersComplete(ctx, id)
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
id := joinedteamchannelmember.NewMeJoinedTeamIdChannelIdMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

payload := joinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
