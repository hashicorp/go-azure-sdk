
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamprimarychannelmessage` Documentation

The `groupteamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamprimarychannelmessage"
```


### Client Initialization

```go
client := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.CreateGroupByIdTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupID("groupIdValue")

payload := groupteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateGroupByIdTeamPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.CreateGroupByIdTeamPrimaryChannelMessageByIdSoftDelete`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.CreateGroupByIdTeamPrimaryChannelMessageByIdSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.CreateGroupByIdTeamPrimaryChannelMessageByIdUndoSoftDelete`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.CreateGroupByIdTeamPrimaryChannelMessageByIdUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.DeleteGroupByIdTeamPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.DeleteGroupByIdTeamPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.GetGroupByIdTeamPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.GetGroupByIdTeamPrimaryChannelMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.GetGroupByIdTeamPrimaryChannelMessageCount`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamPrimaryChannelMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.ListGroupByIdTeamPrimaryChannelMessages`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamPrimaryChannelMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamPrimaryChannelMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.SetGroupByIdTeamPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := groupteamprimarychannelmessage.SetGroupByIdTeamPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.SetGroupByIdTeamPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.UnsetGroupByIdTeamPrimaryChannelMessageByIdReaction`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := groupteamprimarychannelmessage.UnsetGroupByIdTeamPrimaryChannelMessageByIdReactionRequest{
	// ...
}


read, err := client.UnsetGroupByIdTeamPrimaryChannelMessageByIdReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMessageClient.UpdateGroupByIdTeamPrimaryChannelMessageById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := groupteamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateGroupByIdTeamPrimaryChannelMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
