
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/thread` Documentation

The `thread` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/thread"
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


### Example Usage: `ThreadClient.CreateThreadReply`

```go
ctx := context.TODO()
id := thread.NewGroupIdThreadID("groupIdValue", "conversationThreadIdValue")

payload := thread.CreateThreadReplyRequest{
	// ...
}


read, err := client.CreateThreadReply(ctx, id, payload)
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

read, err := client.DeleteThread(ctx, id)
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

read, err := client.GetThread(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadClient.GetThreadCount`

```go
ctx := context.TODO()
id := thread.NewGroupID("groupIdValue")

read, err := client.GetThreadCount(ctx, id)
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

// alternatively `client.ListThreads(ctx, id)` can be used to do batched pagination
items, err := client.ListThreadsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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
