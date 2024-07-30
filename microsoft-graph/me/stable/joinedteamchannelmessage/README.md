
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmessage` Documentation

The `joinedteamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamchannelmessage"
```


### Client Initialization

```go
client := joinedteamchannelmessage.NewJoinedTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelMessageClient.CreateJoinedTeamChannelMessage`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteJoinedTeamChannelMessage(ctx, id)
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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetJoinedTeamChannelMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.GetJoinedTeamChannelMessageCount`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

read, err := client.GetJoinedTeamChannelMessageCount(ctx, id)
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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelID("teamIdValue", "channelIdValue")

// alternatively `client.ListJoinedTeamChannelMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelMessageClient.SetMeJoinedTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessage.SetMeJoinedTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelMessageClient.UnsetMeJoinedTeamChannelMessageReaction`

```go
ctx := context.TODO()
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := joinedteamchannelmessage.UnsetMeJoinedTeamChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamChannelMessageReaction(ctx, id, payload)
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
id := joinedteamchannelmessage.NewMeJoinedTeamIdChannelIdMessageID("teamIdValue", "channelIdValue", "chatMessageIdValue")

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
