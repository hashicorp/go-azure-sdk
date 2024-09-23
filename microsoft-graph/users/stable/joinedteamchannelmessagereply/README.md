
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessagereply` Documentation

The `joinedteamchannelmessagereply` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessagereply"
```


### Client Initialization

```go
client := joinedteamchannelmessagereply.NewJoinedTeamChannelMessageReplyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMessageReplyClient.CreateJoinedTeamChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageID("userId", "teamId", "channelId", "chatMessageId")

payload := joinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamChannelMessageReply(ctx, id, payload, joinedteamchannelmessagereply.DefaultCreateJoinedTeamChannelMessageReplyOperationOptions())
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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.CreateJoinedTeamChannelMessageReplySoftDelete(ctx, id, joinedteamchannelmessagereply.DefaultCreateJoinedTeamChannelMessageReplySoftDeleteOperationOptions())
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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.CreateJoinedTeamChannelMessageReplyUndoSoftDelete(ctx, id, joinedteamchannelmessagereply.DefaultCreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions())
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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageID("userId", "teamId", "channelId", "chatMessageId")

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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageID("userId", "teamId", "channelId", "chatMessageId")

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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

payload := joinedteamchannelmessagereply.SetJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamChannelMessageReplyReaction(ctx, id, payload, joinedteamchannelmessagereply.DefaultSetJoinedTeamChannelMessageReplyReactionOperationOptions())
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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

payload := joinedteamchannelmessagereply.UnsetJoinedTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamChannelMessageReplyReaction(ctx, id, payload, joinedteamchannelmessagereply.DefaultUnsetJoinedTeamChannelMessageReplyReactionOperationOptions())
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
id := joinedteamchannelmessagereply.NewUserIdJoinedTeamIdChannelIdMessageIdReplyID("userId", "teamId", "channelId", "chatMessageId", "chatMessageId1")

payload := joinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamChannelMessageReply(ctx, id, payload, joinedteamchannelmessagereply.DefaultUpdateJoinedTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
