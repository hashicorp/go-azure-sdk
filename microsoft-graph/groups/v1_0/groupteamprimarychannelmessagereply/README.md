
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupteamprimarychannelmessagereply` Documentation

The `groupteamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.CreateGroupByIdTeamPrimaryChannelMessageByIdReply`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := groupteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateGroupByIdTeamPrimaryChannelMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.CreateGroupByIdTeamPrimaryChannelMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateGroupByIdTeamPrimaryChannelMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.CreateGroupByIdTeamPrimaryChannelMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateGroupByIdTeamPrimaryChannelMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.DeleteGroupByIdTeamPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteGroupByIdTeamPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.GetGroupByIdTeamPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetGroupByIdTeamPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.GetGroupByIdTeamPrimaryChannelMessageByIdReplyCount`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.GetGroupByIdTeamPrimaryChannelMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.ListGroupByIdTeamPrimaryChannelMessageByIdReplies`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

// alternatively `client.ListGroupByIdTeamPrimaryChannelMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamPrimaryChannelMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.SetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamprimarychannelmessagereply.SetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.UnsetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamprimarychannelmessagereply.UnsetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetGroupByIdTeamPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageReplyClient.UpdateGroupByIdTeamPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyID("groupIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateGroupByIdTeamPrimaryChannelMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
