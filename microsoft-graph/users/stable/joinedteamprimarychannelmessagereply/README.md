
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessagereply` Documentation

The `joinedteamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessagereply.NewJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.CreateJoinedTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteJoinedTeamPrimaryChannelMessageReply(ctx, id)
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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetJoinedTeamPrimaryChannelMessageReply(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.GetJoinedTeamPrimaryChannelMessageReplyCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamPrimaryChannelMessageReplyCount(ctx, id)
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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

// alternatively `client.ListJoinedTeamPrimaryChannelMessageReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMessageRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.SetUserJoinedTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamprimarychannelmessagereply.SetUserJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetUserJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageReplyClient.UnsetUserJoinedTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := joinedteamprimarychannelmessagereply.UnsetUserJoinedTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetUserJoinedTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessagereply.NewUserIdJoinedTeamIdPrimaryChannelMessageIdReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
