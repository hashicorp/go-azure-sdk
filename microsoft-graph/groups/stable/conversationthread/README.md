
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthread` Documentation

The `conversationthread` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthread"
```


### Client Initialization

```go
client := conversationthread.NewConversationThreadClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadClient.CreateConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationID("groupIdValue", "conversationIdValue")

payload := conversationthread.ConversationThread{
	// ...
}


read, err := client.CreateConversationThread(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.CreateConversationThreadReply`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

payload := conversationthread.CreateConversationThreadReplyRequest{
	// ...
}


read, err := client.CreateConversationThreadReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.DeleteConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.DeleteConversationThread(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.GetConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.GetConversationThread(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.GetConversationThreadCount`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationID("groupIdValue", "conversationIdValue")

read, err := client.GetConversationThreadCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.ListConversationThreads`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationID("groupIdValue", "conversationIdValue")

// alternatively `client.ListConversationThreads(ctx, id)` can be used to do batched pagination
items, err := client.ListConversationThreadsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConversationThreadClient.UpdateConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

payload := conversationthread.ConversationThread{
	// ...
}


read, err := client.UpdateConversationThread(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
