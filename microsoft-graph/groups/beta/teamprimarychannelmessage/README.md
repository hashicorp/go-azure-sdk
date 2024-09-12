
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessage` Documentation

The `teamprimarychannelmessage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessage"
```


### Client Initialization

```go
client := teamprimarychannelmessage.NewTeamPrimaryChannelMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupIdValue")

payload := teamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.CreateTeamPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessageSoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.CreateTeamPrimaryChannelMessageSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.CreateTeamPrimaryChannelMessageUndoSoftDelete`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.CreateTeamPrimaryChannelMessageUndoSoftDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.DeleteTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.DeleteTeamPrimaryChannelMessage(ctx, id, teamprimarychannelmessage.DefaultDeleteTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.GetTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

read, err := client.GetTeamPrimaryChannelMessage(ctx, id, teamprimarychannelmessage.DefaultGetTeamPrimaryChannelMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.GetTeamPrimaryChannelMessagesCount`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupIdValue")

read, err := client.GetTeamPrimaryChannelMessagesCount(ctx, id, teamprimarychannelmessage.DefaultGetTeamPrimaryChannelMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.ListTeamPrimaryChannelMessages`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupID("groupIdValue")

// alternatively `client.ListTeamPrimaryChannelMessages(ctx, id, teamprimarychannelmessage.DefaultListTeamPrimaryChannelMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelMessagesComplete(ctx, id, teamprimarychannelmessage.DefaultListTeamPrimaryChannelMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.SetTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := teamprimarychannelmessage.SetTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.SetTeamPrimaryChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.UnsetTeamPrimaryChannelMessageReaction`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := teamprimarychannelmessage.UnsetTeamPrimaryChannelMessageReactionRequest{
	// ...
}


read, err := client.UnsetTeamPrimaryChannelMessageReaction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMessageClient.UpdateTeamPrimaryChannelMessage`

```go
ctx := context.TODO()
id := teamprimarychannelmessage.NewGroupIdTeamPrimaryChannelMessageID("groupIdValue", "chatMessageIdValue")

payload := teamprimarychannelmessage.ChatMessage{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
