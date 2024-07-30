
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostinreplyto` Documentation

The `conversationthreadpostinreplyto` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostinreplyto"
```


### Client Initialization

```go
client := conversationthreadpostinreplyto.NewConversationThreadPostInReplyToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostInReplyToClient.CreateConversationThreadPostInReplyToForward`

```go
ctx := context.TODO()
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpostinreplyto.CreateConversationThreadPostInReplyToForwardRequest{
	// ...
}


read, err := client.CreateConversationThreadPostInReplyToForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToClient.CreateConversationThreadPostInReplyToReply`

```go
ctx := context.TODO()
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpostinreplyto.CreateConversationThreadPostInReplyToReplyRequest{
	// ...
}


read, err := client.CreateConversationThreadPostInReplyToReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToClient.GetConversationThreadPostInReplyTo`

```go
ctx := context.TODO()
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetConversationThreadPostInReplyTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
