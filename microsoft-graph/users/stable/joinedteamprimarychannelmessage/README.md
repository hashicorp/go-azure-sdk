
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessage` Documentation

The `joinedteamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessage"
```


### Client Initialization

```go
client := joinedteamprimarychannelmessage.NewJoinedTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.CreateJoinedTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.DeleteJoinedTeamPrimaryChannelMessage(ctx, id)
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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamPrimaryChannelMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.GetJoinedTeamPrimaryChannelMessageCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeamPrimaryChannelMessageCount(ctx, id)
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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListJoinedTeamPrimaryChannelMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.SetUserJoinedTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessage.SetUserJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.SetUserJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMessageClient.UnsetUserJoinedTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := joinedteamprimarychannelmessage.UnsetUserJoinedTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetUserJoinedTeamPrimaryChannelMessageReaction(ctx, id, payload)
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
id := joinedteamprimarychannelmessage.NewUserIdJoinedTeamIdPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

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
