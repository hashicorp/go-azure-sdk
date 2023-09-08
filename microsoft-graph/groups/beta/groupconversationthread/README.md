
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthread` Documentation

The `groupconversationthread` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthread"
```


### Client Initialization

```go
client := groupconversationthread.NewGroupConversationThreadClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupConversationThreadClient.CreateGroupByIdConversationByIdThread`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationID("groupIdValue", "conversationIdValue")

payload := groupconversationthread.ConversationThread{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThread(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadClient.CreateGroupByIdConversationByIdThreadByIdReply`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

payload := groupconversationthread.CreateGroupByIdConversationByIdThreadByIdReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadClient.DeleteGroupByIdConversationByIdThreadById`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.DeleteGroupByIdConversationByIdThreadById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadClient.GetGroupByIdConversationByIdThreadById`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

read, err := client.GetGroupByIdConversationByIdThreadById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadClient.GetGroupByIdConversationByIdThreadCount`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationID("groupIdValue", "conversationIdValue")

read, err := client.GetGroupByIdConversationByIdThreadCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadClient.ListGroupByIdConversationByIdThreads`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationID("groupIdValue", "conversationIdValue")

// alternatively `client.ListGroupByIdConversationByIdThreads(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdConversationByIdThreadsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupConversationThreadClient.UpdateGroupByIdConversationByIdThreadById`

```go
ctx := context.TODO()
id := groupconversationthread.NewGroupConversationThreadID("groupIdValue", "conversationIdValue", "conversationThreadIdValue")

payload := groupconversationthread.ConversationThread{
	// ...
}


read, err := client.UpdateGroupByIdConversationByIdThreadById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
