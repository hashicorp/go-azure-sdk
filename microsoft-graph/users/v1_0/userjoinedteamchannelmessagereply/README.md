
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmessagereply` Documentation

The `userjoinedteamchannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamchannelmessagereply"
```


### Client Initialization

```go
client := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReply`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

payload := userjoinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.DeleteUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.GetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.GetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyCount`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.ListUserByIdJoinedTeamByIdChannelByIdMessageByIdReplies`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdChannelByIdMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamchannelmessagereply.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamchannelmessagereply.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamChannelMessageReplyClient.UpdateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyID("userIdValue", "teamIdValue", "channelIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamchannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdChannelByIdMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
