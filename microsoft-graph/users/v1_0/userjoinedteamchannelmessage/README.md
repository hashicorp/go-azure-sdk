
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmessage` Documentation

The `userjoinedteamchannelmessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmessage"
```


### Client Initialization

```go
client := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamChannelMessageClient.CreateUserByIdJoinedTeamByIdChannelByIdMessage`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := userjoinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.DeleteUserByIdJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.GetUserByIdJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.GetUserByIdJoinedTeamByIdChannelByIdMessageCount`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.ListUserByIdJoinedTeamByIdChannelByIdMessages`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelID("userIdValue", "teamIdValue", "channelIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdChannelByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdChannelByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := userjoinedteamchannelmessage.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := userjoinedteamchannelmessage.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageClient.UpdateUserByIdJoinedTeamByIdChannelByIdMessageById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := userjoinedteamchannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdChannelByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
