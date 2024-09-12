
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessagereply` Documentation

The `joinedteamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessagereply.NewJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.CreateJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMessageReply(ctx, id, payload)
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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateJoinedTeamPrimaryChannelMessageReplySoftDelete(ctx, id)
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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateJoinedTeamPrimaryChannelMessageReplyUndoSoftDelete(ctx, id)
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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamprimarychannelmessagereply.SetJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamprimarychannelmessagereply.UnsetJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessagereply.NewMeJoinedTeamIdPrimaryChannelMessageIdReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
