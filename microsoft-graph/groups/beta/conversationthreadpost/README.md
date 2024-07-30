
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpost` Documentation

The `conversationthreadpost` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpost"
```


### Client Initialization

```go
client := conversationthreadpost.NewConversationThreadPostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostClient.CreateConversationThreadPostForward`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpost.CreateConversationThreadPostForwardRequest{
	// ...
}


read, err := client.CreateConversationThreadPostForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostClient.CreateConversationThreadPostReply`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpost.CreateConversationThreadPostReplyRequest{
	// ...
}


read, err := client.CreateConversationThreadPostReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostClient.GetConversationThreadPost`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetConversationThreadPost(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostClient.GetConversationThreadPostCount`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.GetConversationThreadPostCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostClient.ListConversationThreadPosts`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

// alternatively `client.ListConversationThreadPosts(ctx, id)` can be used to do batched pagination
items, err := client.ListConversationThreadPostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConversationThreadPostClient.UpdateConversationThreadPost`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpost.Post{
	// ...
}


read, err := client.UpdateConversationThreadPost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
