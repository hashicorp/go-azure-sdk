
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthread` Documentation

The `groupthread` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthread"
```


### Client Initialization

```go
client := groupthread.NewGroupThreadClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupThreadClient.CreateGroupByIdThread`

```go
ctx := context.TODO()
id := groupthread.NewGroupID("groupIdValue")

payload := groupthread.ConversationThread{
	// ...
}


read, err := client.CreateGroupByIdThread(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadClient.CreateGroupByIdThreadByIdReply`

```go
ctx := context.TODO()
id := groupthread.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

payload := groupthread.CreateGroupByIdThreadByIdReplyRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadClient.DeleteGroupByIdThreadById`

```go
ctx := context.TODO()
id := groupthread.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

read, err := client.DeleteGroupByIdThreadById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadClient.GetGroupByIdThreadById`

```go
ctx := context.TODO()
id := groupthread.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

read, err := client.GetGroupByIdThreadById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadClient.GetGroupByIdThreadCount`

```go
ctx := context.TODO()
id := groupthread.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdThreadCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadClient.ListGroupByIdThreads`

```go
ctx := context.TODO()
id := groupthread.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdThreads(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdThreadsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupThreadClient.UpdateGroupByIdThreadById`

```go
ctx := context.TODO()
id := groupthread.NewGroupThreadID("groupIdValue", "conversationThreadIdValue")

payload := groupthread.ConversationThread{
	// ...
}


read, err := client.UpdateGroupByIdThreadById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
