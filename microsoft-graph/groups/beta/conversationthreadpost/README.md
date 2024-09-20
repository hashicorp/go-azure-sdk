
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpost` Documentation

The `conversationthreadpost` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpost"
```


### Client Initialization

```go
client := conversationthreadpost.NewConversationThreadPostClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostClient.ForwardConversationThreadPost`

```go
ctx := context.TODO()
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpost.ForwardConversationThreadPostRequest{
	// ...
}


read, err := client.ForwardConversationThreadPost(ctx, id, payload, conversationthreadpost.DefaultForwardConversationThreadPostOperationOptions())
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
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

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
id := conversationthreadpost.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

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
id := conversationthreadpost.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

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
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpost.ReplyConversationThreadPostRequest{
	// ...
}


read, err := client.ReplyConversationThreadPost(ctx, id, payload, conversationthreadpost.DefaultReplyConversationThreadPostOperationOptions())
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
id := conversationthreadpost.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpost.Post{
	// ...
}


read, err := client.UpdateConversationThreadPost(ctx, id, payload, conversationthreadpost.DefaultUpdateConversationThreadPostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
