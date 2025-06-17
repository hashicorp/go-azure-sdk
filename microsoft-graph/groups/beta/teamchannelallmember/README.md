
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelallmember` Documentation

The `teamchannelallmember` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelallmember"
```


### Client Initialization

```go
client := teamchannelallmember.NewTeamChannelAllMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelAllMemberClient.AddTeamChannelAllMembers`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelallmember.AddTeamChannelAllMembersRequest{
	// ...
}


// alternatively `client.AddTeamChannelAllMembers(ctx, id, payload, teamchannelallmember.DefaultAddTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddTeamChannelAllMembersComplete(ctx, id, payload, teamchannelallmember.DefaultAddTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelAllMemberClient.CreateTeamChannelAllMember`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelallmember.ConversationMember{
	// ...
}


read, err := client.CreateTeamChannelAllMember(ctx, id, payload, teamchannelallmember.DefaultCreateTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelAllMemberClient.DeleteTeamChannelAllMember`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelIdAllMemberID("groupId", "channelId", "conversationMemberId")

read, err := client.DeleteTeamChannelAllMember(ctx, id, teamchannelallmember.DefaultDeleteTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelAllMemberClient.GetTeamChannelAllMember`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelIdAllMemberID("groupId", "channelId", "conversationMemberId")

read, err := client.GetTeamChannelAllMember(ctx, id, teamchannelallmember.DefaultGetTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelAllMemberClient.GetTeamChannelAllMembersCount`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.GetTeamChannelAllMembersCount(ctx, id, teamchannelallmember.DefaultGetTeamChannelAllMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelAllMemberClient.ListTeamChannelAllMembers`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelID("groupId", "channelId")

// alternatively `client.ListTeamChannelAllMembers(ctx, id, teamchannelallmember.DefaultListTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelAllMembersComplete(ctx, id, teamchannelallmember.DefaultListTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelAllMemberClient.RemoveTeamChannelAllMembers`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelallmember.RemoveTeamChannelAllMembersRequest{
	// ...
}


// alternatively `client.RemoveTeamChannelAllMembers(ctx, id, payload, teamchannelallmember.DefaultRemoveTeamChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveTeamChannelAllMembersComplete(ctx, id, payload, teamchannelallmember.DefaultRemoveTeamChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelAllMemberClient.UpdateTeamChannelAllMember`

```go
ctx := context.TODO()
id := teamchannelallmember.NewGroupIdTeamChannelIdAllMemberID("groupId", "channelId", "conversationMemberId")

payload := teamchannelallmember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamChannelAllMember(ctx, id, payload, teamchannelallmember.DefaultUpdateTeamChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
