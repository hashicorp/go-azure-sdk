
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereply` Documentation

The `teamchannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereply"
```


### Client Initialization

```go
client := teamchannelmessagereply.NewTeamChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMessageReplyClient.CreateTeamChannelMessageReply`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := teamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateTeamChannelMessageReply(ctx, id, payload)
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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateTeamChannelMessageReplySoftDelete(ctx, id)
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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateTeamChannelMessageReplyUndoSoftDelete(ctx, id)
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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteTeamChannelMessageReply(ctx, id, teamchannelmessagereply.DefaultDeleteTeamChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageReplyClient.GetTeamChannelMessageRepliesCount`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

// alternatively `client.ListTeamChannelMessageReplies(ctx, id, teamchannelmessagereply.DefaultListTeamChannelMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelMessageRepliesComplete(ctx, id, teamchannelmessagereply.DefaultListTeamChannelMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageReplyClient.SetTeamChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamchannelmessagereply.SetTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetTeamChannelMessageReplyReaction(ctx, id, payload)
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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamchannelmessagereply.UnsetTeamChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetTeamChannelMessageReplyReaction(ctx, id, payload)
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
id := teamchannelmessagereply.NewGroupIdTeamChannelIdMessageIdReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateTeamChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
