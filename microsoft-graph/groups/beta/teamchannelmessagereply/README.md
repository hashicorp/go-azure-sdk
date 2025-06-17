
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereply` Documentation

The `teamchannelmessagereply` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereply"
```


### Client Initialization

```go
client := teamchannelmessagereply.NewTeamChannelMessageReplyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMessageReplyClient.CreateTeamChannelMessageReply`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateTeamChannelMessageReply(ctx, id, payload, teamchannelmessagereply.DefaultCreateTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.CreateTeamChannelMessageReplySoftDelete`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.CreateTeamChannelMessageReplySoftDelete(ctx, id, teamchannelmessagereply.DefaultCreateTeamChannelMessageReplySoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.CreateTeamChannelMessageReplyUndoSoftDelete`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.CreateTeamChannelMessageReplyUndoSoftDelete(ctx, id, teamchannelmessagereply.DefaultCreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.DeleteTeamChannelMessageReply`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.DeleteTeamChannelMessageReply(ctx, id, teamchannelmessagereply.DefaultDeleteTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.ForwardTeamChannelMessageRepliesToChats`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessagereply.ForwardTeamChannelMessageRepliesToChatsRequest{
	// ...
}


// alternatively `client.ForwardTeamChannelMessageRepliesToChats(ctx, id, payload, teamchannelmessagereply.DefaultForwardTeamChannelMessageRepliesToChatsOperationOptions())` can be used to do batched pagination
items, err := client.ForwardTeamChannelMessageRepliesToChatsComplete(ctx, id, payload, teamchannelmessagereply.DefaultForwardTeamChannelMessageRepliesToChatsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageReplyClient.GetTeamChannelMessageRepliesCount`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

read, err := client.GetTeamChannelMessageRepliesCount(ctx, id, teamchannelmessagereply.DefaultGetTeamChannelMessageRepliesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.GetTeamChannelMessageReply`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

read, err := client.GetTeamChannelMessageReply(ctx, id, teamchannelmessagereply.DefaultGetTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.ListTeamChannelMessageReplies`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

// alternatively `client.ListTeamChannelMessageReplies(ctx, id, teamchannelmessagereply.DefaultListTeamChannelMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelMessageRepliesComplete(ctx, id, teamchannelmessagereply.DefaultListTeamChannelMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageReplyClient.ReplyTeamChannelMessageRepliesWithQuote`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessagereply.ReplyTeamChannelMessageRepliesWithQuoteRequest{
	// ...
}


read, err := client.ReplyTeamChannelMessageRepliesWithQuote(ctx, id, payload, teamchannelmessagereply.DefaultReplyTeamChannelMessageRepliesWithQuoteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.SetTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

payload := teamchannelmessagereply.SetTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetTeamChannelMessageReplyReaction(ctx, id, payload, teamchannelmessagereply.DefaultSetTeamChannelMessageReplyReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.UnsetTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

payload := teamchannelmessagereply.UnsetTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetTeamChannelMessageReplyReaction(ctx, id, payload, teamchannelmessagereply.DefaultUnsetTeamChannelMessageReplyReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.UpdateTeamChannelMessageReply`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupId", "channelId", "chatMessageId", "chatMessageId1")

payload := teamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateTeamChannelMessageReply(ctx, id, payload, teamchannelmessagereply.DefaultUpdateTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
