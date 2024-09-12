
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessage` Documentation

The `joinedteamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessage"
```


### Client Initialization

```go
client := joinedteamchannelmessage.NewJoinedTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMessageClient.CreateJoinedTeamChannelMessage`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := joinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateJoinedTeamChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.CreateJoinedTeamChannelMessageSoftDelete`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateJoinedTeamChannelMessageSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.CreateJoinedTeamChannelMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateJoinedTeamChannelMessageUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.DeleteJoinedTeamChannelMessage`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteJoinedTeamChannelMessage(ctx, id, joinedteamchannelmessage.DefaultDeleteJoinedTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.GetJoinedTeamChannelMessage`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamChannelMessage(ctx, id, joinedteamchannelmessage.DefaultGetJoinedTeamChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.GetJoinedTeamChannelMessagesCount`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.GetJoinedTeamChannelMessagesCount(ctx, id, joinedteamchannelmessage.DefaultGetJoinedTeamChannelMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.ListJoinedTeamChannelMessages`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

// alternatively `client.ListJoinedTeamChannelMessages(ctx, id, joinedteamchannelmessage.DefaultListJoinedTeamChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMessagesComplete(ctx, id, joinedteamchannelmessage.DefaultListJoinedTeamChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMessageClient.SetJoinedTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessage.SetJoinedTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.SetJoinedTeamChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.UnsetJoinedTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessage.UnsetJoinedTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetJoinedTeamChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.UpdateJoinedTeamChannelMessage`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewUserIdJoinedTeamIdChannelIdMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateJoinedTeamChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
