
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessagereply` Documentation

The `teamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessagereply"
```


### Client Initialization

```go
client := teamprimarychannelmessagereply.NewTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.CreateTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := teamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateTeamPrimaryChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.CreateTeamPrimaryChannelMessageReplySoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateTeamPrimaryChannelMessageReplySoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.CreateTeamPrimaryChannelMessageReplyUndoSoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateTeamPrimaryChannelMessageReplyUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.DeleteTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteTeamPrimaryChannelMessageReply(ctx, id, teamprimarychannelmessagereply.DefaultDeleteTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.GetTeamPrimaryChannelMessageRepliesCount`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.GetTeamPrimaryChannelMessageRepliesCount(ctx, id, teamprimarychannelmessagereply.DefaultGetTeamPrimaryChannelMessageRepliesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.GetTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetTeamPrimaryChannelMessageReply(ctx, id, teamprimarychannelmessagereply.DefaultGetTeamPrimaryChannelMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.ListTeamPrimaryChannelMessageReplies`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

// alternatively `client.ListTeamPrimaryChannelMessageReplies(ctx, id, teamprimarychannelmessagereply.DefaultListTeamPrimaryChannelMessageRepliesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelMessageRepliesComplete(ctx, id, teamprimarychannelmessagereply.DefaultListTeamPrimaryChannelMessageRepliesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.SetTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamprimarychannelmessagereply.SetTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.SetTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.UnsetTeamPrimaryChannelMessageReplyReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamprimarychannelmessagereply.UnsetTeamPrimaryChannelMessageReplyReactionRequest{
	// ...
}


read, err := client.UnsetTeamPrimaryChannelMessageReplyReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageReplyClient.UpdateTeamPrimaryChannelMessageReply`

```go
ctx := context.TODO()
id := teamprimarychannelmessagereply.NewGroupIdTeamPrimaryChannelMessageIdReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := teamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
