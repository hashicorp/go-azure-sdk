
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/thread` Documentation

The `thread` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/thread"
```


### Client Initialization

```go
client := thread.NewThreadClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadClient.CreateThread`

```go
ctx := context.TODO()
id := thread.NewGroupID("groupIdValue")

payload := thread.ConversationThread{
	// ...
}


read, err := client.CreateThread(ctx, id, payload)
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
id := thread.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

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
id := thread.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

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
id := thread.NewGroupID("groupIdValue")

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
id := thread.NewGroupID("groupIdValue")

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
id := thread.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

payload := thread.ReplyThreadRequest{
	// ...
}


read, err := client.ReplyThread(ctx, id, payload)
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
id := thread.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

payload := thread.ConversationThread{
	// ...
}


read, err := client.UpdateThread(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
