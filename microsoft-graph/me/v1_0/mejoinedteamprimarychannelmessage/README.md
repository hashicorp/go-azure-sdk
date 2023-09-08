
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmessage` Documentation

The `mejoinedteamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmessage"
```


### Client Initialization

```go
client := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.CreateMeJoinedTeamByIdPrimaryChannelMessage`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.DeleteMeJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.DeleteMeJoinedTeamByIdPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.GetMeJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.GetMeJoinedTeamByIdPrimaryChannelMessageCount`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.ListMeJoinedTeamByIdPrimaryChannelMessages`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.ListMeJoinedTeamByIdPrimaryChannelMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdPrimaryChannelMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := mejoinedteamprimarychannelmessage.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := mejoinedteamprimarychannelmessage.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageClient.UpdateMeJoinedTeamByIdPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := mejoinedteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdPrimaryChannelMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
