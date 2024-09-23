
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplyto` Documentation

The `conversationthreadpostinreplyto` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplyto"
```


### Client Initialization

```go
client := conversationthreadpostinreplyto.NewConversationThreadPostInReplyToClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostInReplyToClient.ForwardConversationThreadPostInReplyTo`

```go
ctx := context.TODO()
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostinreplyto.ForwardConversationThreadPostInReplyToRequest{
	// ...
}


read, err := client.ForwardConversationThreadPostInReplyTo(ctx, id, payload, conversationthreadpostinreplyto.DefaultForwardConversationThreadPostInReplyToOperationOptions())
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
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

read, err := client.GetConversationThreadPostInReplyTo(ctx, id, conversationthreadpostinreplyto.DefaultGetConversationThreadPostInReplyToOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToClient.ReplyConversationThreadPostInReplyTo`

```go
ctx := context.TODO()
id := conversationthreadpostinreplyto.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostinreplyto.ReplyConversationThreadPostInReplyToRequest{
	// ...
}


read, err := client.ReplyConversationThreadPostInReplyTo(ctx, id, payload, conversationthreadpostinreplyto.DefaultReplyConversationThreadPostInReplyToOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
