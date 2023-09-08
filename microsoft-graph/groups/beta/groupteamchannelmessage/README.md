
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamchannelmessage` Documentation

The `groupteamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamchannelmessage"
```


### Client Initialization

```go
client := groupteamchannelmessage.NewGroupTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamChannelMessageClient.CreateGroupByIdTeamChannelByIdMessage`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

payload := groupteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateGroupByIdTeamChannelByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.CreateGroupByIdTeamChannelByIdMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateGroupByIdTeamChannelByIdMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.CreateGroupByIdTeamChannelByIdMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateGroupByIdTeamChannelByIdMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.DeleteGroupByIdTeamChannelByIdMessageById`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteGroupByIdTeamChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.GetGroupByIdTeamChannelByIdMessageById`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetGroupByIdTeamChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.GetGroupByIdTeamChannelByIdMessageCount`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

read, err := client.GetGroupByIdTeamChannelByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.ListGroupByIdTeamChannelByIdMessages`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelID("groupIdValue", "channelIdValue")

// alternatively `client.ListGroupByIdTeamChannelByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamChannelByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamChannelMessageClient.SetGroupByIdTeamChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := groupteamchannelmessage.SetGroupByIdTeamChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.SetGroupByIdTeamChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.UnsetGroupByIdTeamChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := groupteamchannelmessage.UnsetGroupByIdTeamChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetGroupByIdTeamChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamChannelMessageClient.UpdateGroupByIdTeamChannelByIdMessageById`

```go
ctx := context.TODO()
id := groupteamchannelmessage.NewGroupTeamChannelMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := groupteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateGroupByIdTeamChannelByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
