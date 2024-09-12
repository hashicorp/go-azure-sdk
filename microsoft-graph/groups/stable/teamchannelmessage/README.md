
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessage` Documentation

The `teamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessage"
```


### Client Initialization

```go
client := teamchannelmessage.NewTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMessageClient.CreateTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

payload := teamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateTeamChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.CreateTeamChannelMessageSoftDelete`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateTeamChannelMessageSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.CreateTeamChannelMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateTeamChannelMessageUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.DeleteTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteTeamChannelMessage(ctx, id, teamchannelmessage.DefaultDeleteTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.GetTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetTeamChannelMessage(ctx, id, teamchannelmessage.DefaultGetTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.GetTeamChannelMessagesCount`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

read, err := client.GetTeamChannelMessagesCount(ctx, id, teamchannelmessage.DefaultGetTeamChannelMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.ListTeamChannelMessages`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupIdValue", "channelIdValue")

// alternatively `client.ListTeamChannelMessages(ctx, id, teamchannelmessage.DefaultListTeamChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelMessagesComplete(ctx, id, teamchannelmessage.DefaultListTeamChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageClient.SetTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := teamchannelmessage.SetTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.SetTeamChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.UnsetTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := teamchannelmessage.UnsetTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetTeamChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.UpdateTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupIdValue", "channelIdValue", "chatMessageIdValue")

payload := teamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateTeamChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
