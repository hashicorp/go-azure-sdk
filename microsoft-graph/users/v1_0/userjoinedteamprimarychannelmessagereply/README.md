
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmessagereply` Documentation

The `userjoinedteamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReply`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := userjoinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.DeleteUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.GetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.GetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyCount`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.ListUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplies`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdPrimaryChannelMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamprimarychannelmessagereply.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamprimarychannelmessagereply.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageReplyClient.UpdateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyID("userIdValue", "teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := userjoinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
