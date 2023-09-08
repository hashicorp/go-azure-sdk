
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmember` Documentation

The `mejoinedteamchannelmember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmember"
```


### Client Initialization

```go
client := mejoinedteamchannelmember.NewMeJoinedTeamChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamChannelMemberClient.AddMeJoinedTeamByIdChannelByIdMember`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

payload := mejoinedteamchannelmember.AddMeJoinedTeamByIdChannelByIdMemberRequest{
	// ...
}


read, err := client.AddMeJoinedTeamByIdChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.CreateMeJoinedTeamByIdChannelByIdMember`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

payload := mejoinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdChannelByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.DeleteMeJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.DeleteMeJoinedTeamByIdChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.GetMeJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

read, err := client.GetMeJoinedTeamByIdChannelByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.GetMeJoinedTeamByIdChannelByIdMemberCount`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

read, err := client.GetMeJoinedTeamByIdChannelByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.ListMeJoinedTeamByIdChannelByIdMembers`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

// alternatively `client.ListMeJoinedTeamByIdChannelByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdChannelByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamChannelMemberClient.UpdateMeJoinedTeamByIdChannelByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmember.NewMeJoinedTeamChannelMemberID("teamIdValue", "channelIdValue", "conversationMemberIdValue")

payload := mejoinedteamchannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdChannelByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
