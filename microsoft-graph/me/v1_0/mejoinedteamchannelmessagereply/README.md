
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmessagereply` Documentation

The `mejoinedteamchannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamchannelmessagereply"
```


### Client Initialization

```go
client := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.CreateMeJoinedTeamByIdChannelByIdMessageByIdReply`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := mejoinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdChannelByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.CreateMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.CreateMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.DeleteMeJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteMeJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.GetMeJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetMeJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.GetMeJoinedTeamByIdChannelByIdMessageByIdReplyCount`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetMeJoinedTeamByIdChannelByIdMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.ListMeJoinedTeamByIdChannelByIdMessageByIdReplies`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

// alternatively `client.ListMeJoinedTeamByIdChannelByIdMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdChannelByIdMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.SetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamchannelmessagereply.SetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamchannelmessagereply.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamChannelMessageReplyClient.UpdateMeJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
