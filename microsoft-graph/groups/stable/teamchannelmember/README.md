
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmember` Documentation

The `teamchannelmember` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmember"
```


### Client Initialization

```go
client := teamchannelmember.NewTeamChannelMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMemberClient.AddTeamChannelMembers`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmember.AddTeamChannelMembersRequest{
	// ...
}


// alternatively `client.AddTeamChannelMembers(ctx, id, payload, teamchannelmember.DefaultAddTeamChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddTeamChannelMembersComplete(ctx, id, payload, teamchannelmember.DefaultAddTeamChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMemberClient.CreateTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateTeamChannelMember(ctx, id, payload, teamchannelmember.DefaultCreateTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.DeleteTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupId", "channelId", "conversationMemberId")

read, err := client.DeleteTeamChannelMember(ctx, id, teamchannelmember.DefaultDeleteTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.GetTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupId", "channelId", "conversationMemberId")

read, err := client.GetTeamChannelMember(ctx, id, teamchannelmember.DefaultGetTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.GetTeamChannelMembersCount`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.GetTeamChannelMembersCount(ctx, id, teamchannelmember.DefaultGetTeamChannelMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.ListTeamChannelMembers`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupId", "channelId")

// alternatively `client.ListTeamChannelMembers(ctx, id, teamchannelmember.DefaultListTeamChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelMembersComplete(ctx, id, teamchannelmember.DefaultListTeamChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMemberClient.RemoveTeamChannelMembers`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmember.RemoveTeamChannelMembersRequest{
	// ...
}


// alternatively `client.RemoveTeamChannelMembers(ctx, id, payload, teamchannelmember.DefaultRemoveTeamChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveTeamChannelMembersComplete(ctx, id, payload, teamchannelmember.DefaultRemoveTeamChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMemberClient.UpdateTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupId", "channelId", "conversationMemberId")

payload := teamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamChannelMember(ctx, id, payload, teamchannelmember.DefaultUpdateTeamChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
