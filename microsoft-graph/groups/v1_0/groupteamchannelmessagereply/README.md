
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupteamchannelmessagereply` Documentation

The `groupteamchannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupteamchannelmessagereply"
```


### Client Initialization

```go
client := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamChannelMessageReplyClient.CreateGroupByIdTeamChannelByIdMessageByIdReply`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := groupteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateGroupByIdTeamChannelByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.CreateGroupByIdTeamChannelByIdMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateGroupByIdTeamChannelByIdMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.CreateGroupByIdTeamChannelByIdMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateGroupByIdTeamChannelByIdMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.DeleteGroupByIdTeamChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteGroupByIdTeamChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.GetGroupByIdTeamChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetGroupByIdTeamChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.GetGroupByIdTeamChannelByIdMessageByIdReplyCount`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetGroupByIdTeamChannelByIdMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.ListGroupByIdTeamChannelByIdMessageByIdReplies`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

// alternatively `client.ListGroupByIdTeamChannelByIdMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamChannelByIdMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.SetGroupByIdTeamChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamchannelmessagereply.SetGroupByIdTeamChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetGroupByIdTeamChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.UnsetGroupByIdTeamChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamchannelmessagereply.UnsetGroupByIdTeamChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetGroupByIdTeamChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageReplyClient.UpdateGroupByIdTeamChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyID("groupIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := groupteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateGroupByIdTeamChannelByIdMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
