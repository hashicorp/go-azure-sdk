
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmessage` Documentation

The `mejoinedteamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmessage"
```


### Client Initialization

```go
client := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamChannelMessageClient.CreateMeJoinedTeamByIdChannelByIdMessage`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

payload := mejoinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdChannelByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.CreateMeJoinedTeamByIdChannelByIdMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateMeJoinedTeamByIdChannelByIdMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.CreateMeJoinedTeamByIdChannelByIdMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateMeJoinedTeamByIdChannelByIdMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.DeleteMeJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteMeJoinedTeamByIdChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.GetMeJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetMeJoinedTeamByIdChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.GetMeJoinedTeamByIdChannelByIdMessageCount`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

read, err := client.GetMeJoinedTeamByIdChannelByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.ListMeJoinedTeamByIdChannelByIdMessages`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelID("teamIdValue", "channelIdValue")

// alternatively `client.ListMeJoinedTeamByIdChannelByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdChannelByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.SetMeJoinedTeamByIdChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := mejoinedteamchannelmessage.SetMeJoinedTeamByIdChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamByIdChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := mejoinedteamchannelmessage.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageClient.UpdateMeJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := mejoinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdChannelByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
