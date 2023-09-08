
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmessagereply` Documentation

The `mejoinedteamprimarychannelmessagereply` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmessagereply"
```


### Client Initialization

```go
client := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReply`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

payload := mejoinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.DeleteMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.DeleteMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.GetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.GetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyCount`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.ListMeJoinedTeamByIdPrimaryChannelMessageByIdReplies`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageID("teamIdValue", "chatMessageIdValue")

// alternatively `client.ListMeJoinedTeamByIdPrimaryChannelMessageByIdReplies(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdPrimaryChannelMessageByIdRepliesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamprimarychannelmessagereply.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.SetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamprimarychannelmessagereply.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReactionRequest{
	// ...
}


read, err := client.UnsetMeJoinedTeamByIdPrimaryChannelMessageByIdReplyByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMessageReplyClient.UpdateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyID("teamIdValue", "chatMessageIdValue", "chatMessageId1Value")

payload := mejoinedteamprimarychannelmessagereply.ChatMessage{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdPrimaryChannelMessageByIdReplyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
