
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessage` Documentation

The `teamprimarychannelmessage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessage"
```


### Client Initialization

```go
client := teamprimarychannelmessage.NewTeamPrimaryChannelMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupId")

payload := teamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateTeamPrimaryChannelMessage(ctx, id, payload, teamprimarychannelmessage.DefaultCreateTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessageSoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

read, err := client.CreateTeamPrimaryChannelMessageSoftDelete(ctx, id, teamprimarychannelmessage.DefaultCreateTeamPrimaryChannelMessageSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

read, err := client.CreateTeamPrimaryChannelMessageUndoSoftDelete(ctx, id, teamprimarychannelmessage.DefaultCreateTeamPrimaryChannelMessageUndoSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.DeleteTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

read, err := client.DeleteTeamPrimaryChannelMessage(ctx, id, teamprimarychannelmessage.DefaultDeleteTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.ForwardTeamPrimaryChannelMessagesToChats`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupId")

payload := teamprimarychannelmessage.ForwardTeamPrimaryChannelMessagesToChatsRequest{
	// ...
}


// alternatively `client.ForwardTeamPrimaryChannelMessagesToChats(ctx, id, payload, teamprimarychannelmessage.DefaultForwardTeamPrimaryChannelMessagesToChatsOperationOptions())` can be used to do batched pagination
items, err := client.ForwardTeamPrimaryChannelMessagesToChatsComplete(ctx, id, payload, teamprimarychannelmessage.DefaultForwardTeamPrimaryChannelMessagesToChatsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.GetTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

read, err := client.GetTeamPrimaryChannelMessage(ctx, id, teamprimarychannelmessage.DefaultGetTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.GetTeamPrimaryChannelMessagesCount`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupId")

read, err := client.GetTeamPrimaryChannelMessagesCount(ctx, id, teamprimarychannelmessage.DefaultGetTeamPrimaryChannelMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.ListTeamPrimaryChannelMessages`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupId")

// alternatively `client.ListTeamPrimaryChannelMessages(ctx, id, teamprimarychannelmessage.DefaultListTeamPrimaryChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelMessagesComplete(ctx, id, teamprimarychannelmessage.DefaultListTeamPrimaryChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.ReplyTeamPrimaryChannelMessagesWithQuote`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupId")

payload := teamprimarychannelmessage.ReplyTeamPrimaryChannelMessagesWithQuoteRequest{
	// ...
}


read, err := client.ReplyTeamPrimaryChannelMessagesWithQuote(ctx, id, payload, teamprimarychannelmessage.DefaultReplyTeamPrimaryChannelMessagesWithQuoteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.SetTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

payload := teamprimarychannelmessage.SetTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.SetTeamPrimaryChannelMessageReaction(ctx, id, payload, teamprimarychannelmessage.DefaultSetTeamPrimaryChannelMessageReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.UnsetTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

payload := teamprimarychannelmessage.UnsetTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetTeamPrimaryChannelMessageReaction(ctx, id, payload, teamprimarychannelmessage.DefaultUnsetTeamPrimaryChannelMessageReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.UpdateTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupId", "chatMessageId")

payload := teamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelMessage(ctx, id, payload, teamprimarychannelmessage.DefaultUpdateTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
