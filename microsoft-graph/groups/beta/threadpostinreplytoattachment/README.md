
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoattachment` Documentation

The `threadpostinreplytoattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoattachment"
```


### Client Initialization

```go
client := threadpostinreplytoattachment.NewThreadPostInReplyToAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.CreateThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpostinreplytoattachment.Attachment{
	// ...
}


read, err := client.CreateThreadPostInReplyToAttachment(ctx, id, payload, threadpostinreplytoattachment.DefaultCreateThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.CreateThreadPostInReplyToAttachmentsUploadSession`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpostinreplytoattachment.CreateThreadPostInReplyToAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateThreadPostInReplyToAttachmentsUploadSession(ctx, id, payload, threadpostinreplytoattachment.DefaultCreateThreadPostInReplyToAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.DeleteThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationThreadId", "postId", "attachmentId")

read, err := client.DeleteThreadPostInReplyToAttachment(ctx, id, threadpostinreplytoattachment.DefaultDeleteThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.GetThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationThreadId", "postId", "attachmentId")

read, err := client.GetThreadPostInReplyToAttachment(ctx, id, threadpostinreplytoattachment.DefaultGetThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.GetThreadPostInReplyToAttachmentsCount`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

read, err := client.GetThreadPostInReplyToAttachmentsCount(ctx, id, threadpostinreplytoattachment.DefaultGetThreadPostInReplyToAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.ListThreadPostInReplyToAttachments`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

// alternatively `client.ListThreadPostInReplyToAttachments(ctx, id, threadpostinreplytoattachment.DefaultListThreadPostInReplyToAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListThreadPostInReplyToAttachmentsComplete(ctx, id, threadpostinreplytoattachment.DefaultListThreadPostInReplyToAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
