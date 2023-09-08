
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupthreadpostinreplytoattachment` Documentation

The `groupthreadpostinreplytoattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupthreadpostinreplytoattachment"
```


### Client Initialization

```go
client := groupthreadpostinreplytoattachment.NewGroupThreadPostInReplyToAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.CreateGroupByIdThreadByIdPostByIdInReplyToAttachment`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostinreplytoattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdInReplyToAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.CreateGroupByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupthreadpostinreplytoattachment.CreateGroupByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.DeleteGroupByIdThreadByIdPostByIdInReplyToAttachmentById`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostInReplyToAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdThreadByIdPostByIdInReplyToAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.GetGroupByIdThreadByIdPostByIdInReplyToAttachmentById`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostInReplyToAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdThreadByIdPostByIdInReplyToAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.GetGroupByIdThreadByIdPostByIdInReplyToAttachmentCount`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdThreadByIdPostByIdInReplyToAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupThreadPostInReplyToAttachmentClient.ListGroupByIdThreadByIdPostByIdInReplyToAttachments`

```go
ctx := context.TODO()
id := groupthreadpostinreplytoattachment.NewGroupThreadPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListGroupByIdThreadByIdPostByIdInReplyToAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdThreadByIdPostByIdInReplyToAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
