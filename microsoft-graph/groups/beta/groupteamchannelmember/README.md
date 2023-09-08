
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamchannelmember` Documentation

The `groupteamchannelmember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamchannelmember"
```


### Client Initialization

```go
client := groupteamchannelmember.NewGroupTeamChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamChannelMemberClient.AddGroupByIdTeamChannelByIdMember`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

payload := groupteamchannelmember.AddGroupByIdTeamChannelByIdMemberRequest{
	// ...
}


read, err := client.AddGroupByIdTeamChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMemberClient.CreateGroupByIdTeamChannelByIdMember`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

payload := groupteamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateGroupByIdTeamChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMemberClient.DeleteGroupByIdTeamChannelByIdMemberById`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.DeleteGroupByIdTeamChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMemberClient.GetGroupByIdTeamChannelByIdMemberById`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.GetGroupByIdTeamChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMemberClient.GetGroupByIdTeamChannelByIdMemberCount`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

read, err := client.GetGroupByIdTeamChannelByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMemberClient.ListGroupByIdTeamChannelByIdMembers`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

// alternatively `client.ListGroupByIdTeamChannelByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamChannelByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamChannelMemberClient.UpdateGroupByIdTeamChannelByIdMemberById`

```go
ctx := context.TODO()
id := groupteamchannelmember.NewGroupTeamChannelMemberID("groupIdValue", "channelIdValue", "conversationMemberIdValue")

payload := groupteamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateGroupByIdTeamChannelByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
