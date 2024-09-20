
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/thread` Documentation

The `thread` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/thread"
```


### Client Initialization

```go
client := thread.NewThreadClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadClient.CreateThread`

```go
ctx := context.TODO()
id := thread.NewGroupID("groupId")

payload := thread.ConversationThread{
	// ...
}


read, err := client.CreateThread(ctx, id, payload, thread.DefaultCreateThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.DeleteThread`

```go
ctx := context.TODO()
id := thread.NewGroupIdThreadID("groupId", "conversationThreadId")

read, err := client.DeleteThread(ctx, id, thread.DefaultDeleteThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.GetThread`

```go
ctx := context.TODO()
id := thread.NewGroupIdThreadID("groupId", "conversationThreadId")

read, err := client.GetThread(ctx, id, thread.DefaultGetThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.GetThreadsCount`

```go
ctx := context.TODO()
id := thread.NewGroupID("groupId")

read, err := client.GetThreadsCount(ctx, id, thread.DefaultGetThreadsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.ListThreads`

```go
ctx := context.TODO()
id := thread.NewGroupID("groupId")

// alternatively `client.ListThreads(ctx, id, thread.DefaultListThreadsOperationOptions())` can be used to do batched pagination
items, err := client.ListThreadsComplete(ctx, id, thread.DefaultListThreadsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ThreadClient.ReplyThread`

```go
ctx := context.TODO()
id := thread.NewGroupIdThreadID("groupId", "conversationThreadId")

payload := thread.ReplyThreadRequest{
	// ...
}


read, err := client.ReplyThread(ctx, id, payload, thread.DefaultReplyThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.UpdateThread`

```go
ctx := context.TODO()
id := thread.NewGroupIdThreadID("groupId", "conversationThreadId")

payload := thread.ConversationThread{
	// ...
}


read, err := client.UpdateThread(ctx, id, payload, thread.DefaultUpdateThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
