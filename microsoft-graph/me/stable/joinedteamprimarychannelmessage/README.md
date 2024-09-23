
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessage` Documentation

The `joinedteamprimarychannelmessage` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessage"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessage.NewJoinedTeamPrimaryChannelMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.CreateJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamId")

payload := joinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMessage(ctx, id, payload, joinedteamprimarychannelmessage.DefaultCreateJoinedTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.CreateJoinedTeamPrimaryChannelMessageSoftDelete`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

read, err := client.CreateJoinedTeamPrimaryChannelMessageSoftDelete(ctx, id, joinedteamprimarychannelmessage.DefaultCreateJoinedTeamPrimaryChannelMessageSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.CreateJoinedTeamPrimaryChannelMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

read, err := client.CreateJoinedTeamPrimaryChannelMessageUndoSoftDelete(ctx, id, joinedteamprimarychannelmessage.DefaultCreateJoinedTeamPrimaryChannelMessageUndoSoftDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.DeleteJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

read, err := client.DeleteJoinedTeamPrimaryChannelMessage(ctx, id, joinedteamprimarychannelmessage.DefaultDeleteJoinedTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.GetJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

read, err := client.GetJoinedTeamPrimaryChannelMessage(ctx, id, joinedteamprimarychannelmessage.DefaultGetJoinedTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.GetJoinedTeamPrimaryChannelMessagesCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamId")

read, err := client.GetJoinedTeamPrimaryChannelMessagesCount(ctx, id, joinedteamprimarychannelmessage.DefaultGetJoinedTeamPrimaryChannelMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.ListJoinedTeamPrimaryChannelMessages`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamId")

// alternatively `client.ListJoinedTeamPrimaryChannelMessages(ctx, id, joinedteamprimarychannelmessage.DefaultListJoinedTeamPrimaryChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMessagesComplete(ctx, id, joinedteamprimarychannelmessage.DefaultListJoinedTeamPrimaryChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.SetJoinedTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

payload := joinedteamprimarychannelmessage.SetJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload, joinedteamprimarychannelmessage.DefaultSetJoinedTeamPrimaryChannelMessageReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.UnsetJoinedTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

payload := joinedteamprimarychannelmessage.UnsetJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload, joinedteamprimarychannelmessage.DefaultUnsetJoinedTeamPrimaryChannelMessageReactionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.UpdateJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamId", "chatMessageId")

payload := joinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMessage(ctx, id, payload, joinedteamprimarychannelmessage.DefaultUpdateJoinedTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
