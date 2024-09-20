
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthread` Documentation

The `conversationthread` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthread"
```


### Client Initialization

```go
client := conversationthread.NewConversationThreadClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadClient.CreateConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationID("groupId", "conversationId")

payload := conversationthread.ConversationThread{
	// ...
}


read, err := client.CreateConversationThread(ctx, id, payload, conversationthread.DefaultCreateConversationThreadOperationOptions())
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
id := conversationthread.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

read, err := client.DeleteConversationThread(ctx, id, conversationthread.DefaultDeleteConversationThreadOperationOptions())
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
id := conversationthread.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

read, err := client.GetConversationThread(ctx, id, conversationthread.DefaultGetConversationThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.GetConversationThreadsCount`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationID("groupId", "conversationId")

read, err := client.GetConversationThreadsCount(ctx, id, conversationthread.DefaultGetConversationThreadsCountOperationOptions())
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
id := conversationthread.NewGroupIdConversationID("groupId", "conversationId")

// alternatively `client.ListConversationThreads(ctx, id, conversationthread.DefaultListConversationThreadsOperationOptions())` can be used to do batched pagination
items, err := client.ListConversationThreadsComplete(ctx, id, conversationthread.DefaultListConversationThreadsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConversationThreadClient.ReplyConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

payload := conversationthread.ReplyConversationThreadRequest{
	// ...
}


read, err := client.ReplyConversationThread(ctx, id, payload, conversationthread.DefaultReplyConversationThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadClient.UpdateConversationThread`

```go
ctx := context.TODO()
id := conversationthread.NewGroupIdConversationIdThreadID("groupId", "conversationId", "conversationThreadId")

payload := conversationthread.ConversationThread{
	// ...
}


read, err := client.UpdateConversationThread(ctx, id, payload, conversationthread.DefaultUpdateConversationThreadOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
