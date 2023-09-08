
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthreadpostattachment` Documentation

The `groupthreadpostattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupthreadpostattachment"
```


### Client Initialization

```go
client := groupthreadpostattachment.NewGroupThreadPostAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupThreadPostAttachmentClient.CreateGroupByIdThreadByIdPostByIdAttachment`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostAttachmentClient.CreateGroupByIdThreadByIdPostByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostattachment.CreateGroupByIdThreadByIdPostByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostAttachmentClient.DeleteGroupByIdThreadByIdPostByIdAttachmentById`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdThreadByIdPostByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostAttachmentClient.GetGroupByIdThreadByIdPostByIdAttachmentById`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdThreadByIdPostByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostAttachmentClient.GetGroupByIdThreadByIdPostByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdThreadByIdPostByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostAttachmentClient.ListGroupByIdThreadByIdPostByIdAttachments`

```go
ctx := context.TODO()
id := groupthreadpostattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListGroupByIdThreadByIdPostByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdThreadByIdPostByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
