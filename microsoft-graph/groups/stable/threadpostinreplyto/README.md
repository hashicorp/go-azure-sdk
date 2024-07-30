
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostinreplyto` Documentation

The `threadpostinreplyto` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostinreplyto"
```


### Client Initialization

```go
client := threadpostinreplyto.NewThreadPostInReplyToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostInReplyToClient.CreateThreadPostInReplyToForward`

```go
ctx := context.TODO()
id := threadpostinreplyto.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplyto.CreateThreadPostInReplyToForwardRequest{
	// ...
}


read, err := client.CreateThreadPostInReplyToForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToClient.CreateThreadPostInReplyToReply`

```go
ctx := context.TODO()
id := threadpostinreplyto.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplyto.CreateThreadPostInReplyToReplyRequest{
	// ...
}


read, err := client.CreateThreadPostInReplyToReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToClient.GetThreadPostInReplyTo`

```go
ctx := context.TODO()
id := threadpostinreplyto.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetThreadPostInReplyTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
