
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpost` Documentation

The `threadpost` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpost"
```


### Client Initialization

```go
client := threadpost.NewThreadPostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostClient.CreateThreadPostForward`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpost.CreateThreadPostForwardRequest{
	// ...
}


read, err := client.CreateThreadPostForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostClient.CreateThreadPostReply`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpost.CreateThreadPostReplyRequest{
	// ...
}


read, err := client.CreateThreadPostReply(ctx, id, payload)
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
id := threadpost.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetThreadPost(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostClient.GetThreadPostCount`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

read, err := client.GetThreadPostCount(ctx, id)
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
id := threadpost.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

// alternatively `client.ListThreadPosts(ctx, id)` can be used to do batched pagination
items, err := client.ListThreadPostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ThreadPostClient.UpdateThreadPost`

```go
ctx := context.TODO()
id := threadpost.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpost.Post{
	// ...
}


read, err := client.UpdateThreadPost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
