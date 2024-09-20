
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpost` Documentation

The `threadpost` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpost"
```


### Client Initialization

```go
client := threadpost.NewThreadPostClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostClient.ForwardThreadPost`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpost.ForwardThreadPostRequest{
	// ...
}


read, err := client.ForwardThreadPost(ctx, id, payload, threadpost.DefaultForwardThreadPostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostClient.GetThreadPost`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

read, err := client.GetThreadPost(ctx, id, threadpost.DefaultGetThreadPostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostClient.GetThreadPostsCount`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadID("groupId", "conversationThreadId")

read, err := client.GetThreadPostsCount(ctx, id, threadpost.DefaultGetThreadPostsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostClient.ListThreadPosts`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadID("groupId", "conversationThreadId")

// alternatively `client.ListThreadPosts(ctx, id, threadpost.DefaultListThreadPostsOperationOptions())` can be used to do batched pagination
items, err := client.ListThreadPostsComplete(ctx, id, threadpost.DefaultListThreadPostsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ThreadPostClient.ReplyThreadPost`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpost.ReplyThreadPostRequest{
	// ...
}


read, err := client.ReplyThreadPost(ctx, id, payload, threadpost.DefaultReplyThreadPostOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
