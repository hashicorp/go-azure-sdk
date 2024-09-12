
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplyto` Documentation

The `threadpostinreplyto` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplyto"
```


### Client Initialization

```go
client := threadpostinreplyto.NewThreadPostInReplyToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostInReplyToClient.ForwardThreadPostInReplyTo`

```go
ctx := context.TODO()
id := threadpostinreplyto.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplyto.ForwardThreadPostInReplyToRequest{
	// ...
}


read, err := client.ForwardThreadPostInReplyTo(ctx, id, payload)
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

read, err := client.GetThreadPostInReplyTo(ctx, id, threadpostinreplyto.DefaultGetThreadPostInReplyToOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToClient.ReplyThreadPostInReplyTo`

```go
ctx := context.TODO()
id := threadpostinreplyto.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplyto.ReplyThreadPostInReplyToRequest{
	// ...
}


read, err := client.ReplyThreadPostInReplyTo(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
