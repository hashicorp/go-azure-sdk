
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmessagereply` Documentation

The `joinedteamchannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmessagereply"
```


### Client Initialization

```go
client := joinedteamchannelmessagereply.NewJoinedTeamChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.CreateJoinedTeamChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.CreateJoinedTeamChannelMessageReplySoftDelete`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateJoinedTeamChannelMessageReplySoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.CreateJoinedTeamChannelMessageReplyUndoSoftDelete`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateJoinedTeamChannelMessageReplyUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.DeleteJoinedTeamChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteJoinedTeamChannelMessageReply(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.GetJoinedTeamChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetJoinedTeamChannelMessageReply(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.GetJoinedTeamChannelMessageReplyCount`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamChannelMessageReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.ListJoinedTeamChannelMessageReplies`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

// alternatively `client.ListJoinedTeamChannelMessageReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMessageRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.SetMeJoinedTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamchannelmessagereply.SetMeJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.UnsetMeJoinedTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamchannelmessagereply.UnsetMeJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.UpdateJoinedTeamChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
