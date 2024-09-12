
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

read, err := client.DeleteJoinedTeamChannelMessageReply(ctx, id, joinedteamchannelmessagereply.DefaultDeleteJoinedTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.GetJoinedTeamChannelMessageRepliesCount`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamChannelMessageRepliesCount(ctx, id, joinedteamchannelmessagereply.DefaultGetJoinedTeamChannelMessageRepliesCountOperationOptions())
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

read, err := client.GetJoinedTeamChannelMessageReply(ctx, id, joinedteamchannelmessagereply.DefaultGetJoinedTeamChannelMessageReplyOperationOptions())
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

// alternatively `client.ListJoinedTeamChannelMessageReplies(ctx, id, joinedteamchannelmessagereply.DefaultListJoinedTeamChannelMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMessageRepliesComplete(ctx, id, joinedteamchannelmessagereply.DefaultListJoinedTeamChannelMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.SetJoinedTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamchannelmessagereply.SetJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.UnsetJoinedTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewMeJoinedTeamIdChannelIdMessageIdReplyID("teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamchannelmessagereply.UnsetJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamChannelMessageReplyReaction(ctx, id, payload)
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
