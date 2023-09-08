
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupthreadpostinreplyto` Documentation

The `groupthreadpostinreplyto` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupthreadpostinreplyto"
```


### Client Initialization

```go
client := groupthreadpostinreplyto.NewGroupThreadPostInReplyToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupThreadPostInReplyToClient.CreateGroupByIdThreadByIdPostByIdInReplyToForward`

```go
ctx := context.TODO()
id := groupthreadpostinreplyto.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostinreplyto.CreateGroupByIdThreadByIdPostByIdInReplyToForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdInReplyToForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToClient.CreateGroupByIdThreadByIdPostByIdInReplyToReply`

```go
ctx := context.TODO()
id := groupthreadpostinreplyto.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostinreplyto.CreateGroupByIdThreadByIdPostByIdInReplyToReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdInReplyToReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToClient.GetGroupByIdThreadByIdPostByIdInReplyTo`

```go
ctx := context.TODO()
id := groupthreadpostinreplyto.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdThreadByIdPostByIdInReplyTo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
