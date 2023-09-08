
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmessage` Documentation

The `userjoinedteamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamprimarychannelmessage"
```


### Client Initialization

```go
client := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessage`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdPrimaryChannelMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.DeleteUserByIdJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.GetUserByIdJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.GetUserByIdJoinedTeamByIdPrimaryChannelMessageCount`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPrimaryChannelMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.ListUserByIdJoinedTeamByIdPrimaryChannelMessages`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdPrimaryChannelMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdPrimaryChannelMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := userjoinedteamprimarychannelmessage.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.SetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := userjoinedteamprimarychannelmessage.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetUserByIdJoinedTeamByIdPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPrimaryChannelMessageClient.UpdateUserByIdJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageID("userIdValue", "teamIdValue", "chatMessageIdValue")

payload := userjoinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdPrimaryChannelMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
