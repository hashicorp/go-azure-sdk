
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthreadpost` Documentation

The `groupthreadpost` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthreadpost"
```


### Client Initialization

```go
client := groupthreadpost.NewGroupThreadPostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupThreadPostClient.CreateGroupByIdThreadByIdPostByIdForward`

```go
ctx := context.TODO()
id := groupthreadpost.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpost.CreateGroupByIdThreadByIdPostByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostClient.CreateGroupByIdThreadByIdPostByIdReply`

```go
ctx := context.TODO()
id := groupthreadpost.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpost.CreateGroupByIdThreadByIdPostByIdReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostClient.GetGroupByIdThreadByIdPostById`

```go
ctx := context.TODO()
id := groupthreadpost.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdThreadByIdPostById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostClient.GetGroupByIdThreadByIdPostCount`

```go
ctx := context.TODO()
id := groupthreadpost.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

read, err := client.GetGroupByIdThreadByIdPostCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostClient.ListGroupByIdThreadByIdPosts`

```go
ctx := context.TODO()
id := groupthreadpost.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

// alternatively `client.ListGroupByIdThreadByIdPosts(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdThreadByIdPostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
