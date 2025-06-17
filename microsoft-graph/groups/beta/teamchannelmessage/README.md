
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessage` Documentation

The `teamchannelmessage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessage"
```


### Client Initialization

```go
client := teamchannelmessage.NewTeamChannelMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelMessageClient.CreateTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateTeamChannelMessage(ctx, id, payload, teamchannelmessage.DefaultCreateTeamChannelMessageOperationOptions())
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
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

read, err := client.CreateTeamChannelMessageSoftDelete(ctx, id, teamchannelmessage.DefaultCreateTeamChannelMessageSoftDeleteOperationOptions())
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
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

read, err := client.CreateTeamChannelMessageUndoSoftDelete(ctx, id, teamchannelmessage.DefaultCreateTeamChannelMessageUndoSoftDeleteOperationOptions())
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
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

read, err := client.DeleteTeamChannelMessage(ctx, id, teamchannelmessage.DefaultDeleteTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.ForwardTeamChannelMessagesToChats`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmessage.ForwardTeamChannelMessagesToChatsRequest{
	// ...
}


// alternatively `client.ForwardTeamChannelMessagesToChats(ctx, id, payload, teamchannelmessage.DefaultForwardTeamChannelMessagesToChatsOperationOptions())` can be used to do batched pagination
items, err := client.ForwardTeamChannelMessagesToChatsComplete(ctx, id, payload, teamchannelmessage.DefaultForwardTeamChannelMessagesToChatsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageClient.GetTeamChannelMessage`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

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
id := teamchannelmessage.NewGroupIdTeamChannelID("groupId", "channelId")

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
id := teamchannelmessage.NewGroupIdTeamChannelID("groupId", "channelId")

// alternatively `client.ListTeamChannelMessages(ctx, id, teamchannelmessage.DefaultListTeamChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelMessagesComplete(ctx, id, teamchannelmessage.DefaultListTeamChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelMessageClient.ReplyTeamChannelMessagesWithQuote`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelmessage.ReplyTeamChannelMessagesWithQuoteRequest{
	// ...
}


read, err := client.ReplyTeamChannelMessagesWithQuote(ctx, id, payload, teamchannelmessage.DefaultReplyTeamChannelMessagesWithQuoteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelMessageClient.SetTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessage.SetTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.SetTeamChannelMessageReaction(ctx, id, payload, teamchannelmessage.DefaultSetTeamChannelMessageReactionOperationOptions())
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
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessage.UnsetTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetTeamChannelMessageReaction(ctx, id, payload, teamchannelmessage.DefaultUnsetTeamChannelMessageReactionOperationOptions())
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
id := teamchannelmessage.NewGroupIdTeamChannelIdMessageID("groupId", "channelId", "chatMessageId")

payload := teamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateTeamChannelMessage(ctx, id, payload, teamchannelmessage.DefaultUpdateTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
