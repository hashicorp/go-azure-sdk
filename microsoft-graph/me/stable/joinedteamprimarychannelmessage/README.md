
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessage` Documentation

The `joinedteamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamprimarychannelmessage"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessage.NewJoinedTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.CreateJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

payload := joinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMessage(ctx, id, payload)
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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.CreateJoinedTeamPrimaryChannelMessageSoftDelete(ctx, id)
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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.CreateJoinedTeamPrimaryChannelMessageUndoSoftDelete(ctx, id)
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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessage.SetJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessage.UnsetJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessage.NewMeJoinedTeamIdPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
