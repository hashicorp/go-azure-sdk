
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpostinreplyto` Documentation

The `groupconversationthreadpostinreplyto` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpostinreplyto"
```


### Client Initialization

```go
client := groupconversationthreadpostinreplyto.NewGroupConversationThreadPostInReplyToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupConversationThreadPostInReplyToClient.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToForward`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplyto.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostinreplyto.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToClient.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToReply`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplyto.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostinreplyto.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToClient.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyTo`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplyto.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
