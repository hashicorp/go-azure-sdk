
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmember` Documentation

The `teamchannelmember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmember"
```


### Client Initialization

```go
client := teamchannelmember.NewTeamChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMemberClient.AddGroupTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

payload := teamchannelmember.AddGroupTeamChannelMemberRequest{
	// ...
}


read, err := client.AddGroupTeamChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.CreateTeamChannelMember`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

payload := teamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateTeamChannelMember(ctx, id, payload)
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
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.DeleteTeamChannelMember(ctx, id)
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
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.GetTeamChannelMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMemberClient.GetTeamChannelMemberCount`

```go
ctx := context.TODO()
id := teamchannelmember.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

read, err := client.GetTeamChannelMemberCount(ctx, id)
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
id := teamchannelmember.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

// alternatively `client.ListTeamChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListTeamChannelMembersComplete(ctx, id)
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
id := teamchannelmember.NewGroupIdTeamChannelIdMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

payload := teamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
