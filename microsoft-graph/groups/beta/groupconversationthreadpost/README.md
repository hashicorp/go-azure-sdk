
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpost` Documentation

The `groupconversationthreadpost` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpost"
```


### Client Initialization

```go
client := groupconversationthreadpost.NewGroupConversationThreadPostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupConversationThreadPostClient.CreateGroupByIdConversationByIdThreadByIdPostByIdForward`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpost.CreateGroupByIdConversationByIdThreadByIdPostByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostClient.CreateGroupByIdConversationByIdThreadByIdPostByIdReply`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpost.CreateGroupByIdConversationByIdThreadByIdPostByIdReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostClient.GetGroupByIdConversationByIdThreadByIdPostById`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostClient.GetGroupByIdConversationByIdThreadByIdPostCount`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostClient.ListGroupByIdConversationByIdThreadByIdPosts`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

// alternatively `client.ListGroupByIdConversationByIdThreadByIdPosts(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdConversationByIdThreadByIdPostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupConversationThreadPostClient.UpdateGroupByIdConversationByIdThreadByIdPostById`

```go
ctx := context.TODO()
id := groupconversationthreadpost.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpost.Post{
	// ...
}


read, err := client.UpdateGroupByIdConversationByIdThreadByIdPostById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
