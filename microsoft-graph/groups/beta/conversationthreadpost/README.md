
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


### Example Usage: `ConversationThreadPostClient.ForwardConversationThreadPost`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpost.ForwardConversationThreadPostRequest{
	// ...
}


read, err := client.ForwardConversationThreadPost(ctx, id, payload)
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

read, err := client.GetConversationThreadPost(ctx, id, conversationthreadpost.DefaultGetConversationThreadPostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostClient.GetConversationThreadPostsCount`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.GetConversationThreadPostsCount(ctx, id, conversationthreadpost.DefaultGetConversationThreadPostsCountOperationOptions())
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

// alternatively `client.ListConversationThreadPosts(ctx, id, conversationthreadpost.DefaultListConversationThreadPostsOperationOptions())` can be used to do batched pagination
items, err := client.ListConversationThreadPostsComplete(ctx, id, conversationthreadpost.DefaultListConversationThreadPostsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConversationThreadPostClient.ReplyConversationThreadPost`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpost.ReplyConversationThreadPostRequest{
	// ...
}


read, err := client.ReplyConversationThreadPost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
