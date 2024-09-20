
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessagereply` Documentation

The `joinedteamprimarychannelmessagereply` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessagereply.NewJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.CreateJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

payload := joinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMessageReply(ctx, id, payload, joinedteamprimarychannelmessagereply.DefaultCreateJoinedTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.CreateJoinedTeamPrimaryChannelMessageReplySoftDelete`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

read, err := client.CreateJoinedTeamPrimaryChannelMessageReplySoftDelete(ctx, id, joinedteamprimarychannelmessagereply.DefaultCreateJoinedTeamPrimaryChannelMessageReplySoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.CreateJoinedTeamPrimaryChannelMessageReplyUndoSoftDelete`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

read, err := client.CreateJoinedTeamPrimaryChannelMessageReplyUndoSoftDelete(ctx, id, joinedteamprimarychannelmessagereply.DefaultCreateJoinedTeamPrimaryChannelMessageReplyUndoSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.DeleteJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

read, err := client.DeleteJoinedTeamPrimaryChannelMessageReply(ctx, id, joinedteamprimarychannelmessagereply.DefaultDeleteJoinedTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.GetJoinedTeamPrimaryChannelMessageRepliesCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

read, err := client.GetJoinedTeamPrimaryChannelMessageRepliesCount(ctx, id, joinedteamprimarychannelmessagereply.DefaultGetJoinedTeamPrimaryChannelMessageRepliesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.GetJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

read, err := client.GetJoinedTeamPrimaryChannelMessageReply(ctx, id, joinedteamprimarychannelmessagereply.DefaultGetJoinedTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.ListJoinedTeamPrimaryChannelMessageReplies`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

// alternatively `client.ListJoinedTeamPrimaryChannelMessageReplies(ctx, id, joinedteamprimarychannelmessagereply.DefaultListJoinedTeamPrimaryChannelMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMessageRepliesComplete(ctx, id, joinedteamprimarychannelmessagereply.DefaultListJoinedTeamPrimaryChannelMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.SetJoinedTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

payload := joinedteamprimarychannelmessagereply.SetJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload, joinedteamprimarychannelmessagereply.DefaultSetJoinedTeamPrimaryChannelMessageReplyReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.UnsetJoinedTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

payload := joinedteamprimarychannelmessagereply.UnsetJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload, joinedteamprimarychannelmessagereply.DefaultUnsetJoinedTeamPrimaryChannelMessageReplyReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.UpdateJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamId", "chatMessageId", "chatMessageId1")

payload := joinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMessageReply(ctx, id, payload, joinedteamprimarychannelmessagereply.DefaultUpdateJoinedTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
