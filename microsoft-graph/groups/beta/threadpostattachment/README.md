
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostattachment` Documentation

The `threadpostattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostattachment"
```


### Client Initialization

```go
client := threadpostattachment.NewThreadPostAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostAttachmentClient.CreateThreadPostAttachment`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpostattachment.Attachment{
	// ...
}


read, err := client.CreateThreadPostAttachment(ctx, id, payload, threadpostattachment.DefaultCreateThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostAttachmentClient.CreateThreadPostAttachmentsUploadSession`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

payload := threadpostattachment.CreateThreadPostAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateThreadPostAttachmentsUploadSession(ctx, id, payload, threadpostattachment.DefaultCreateThreadPostAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostAttachmentClient.DeleteThreadPostAttachment`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostIdAttachmentID("groupId", "conversationThreadId", "postId", "attachmentId")

read, err := client.DeleteThreadPostAttachment(ctx, id, threadpostattachment.DefaultDeleteThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostAttachmentClient.GetThreadPostAttachment`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostIdAttachmentID("groupId", "conversationThreadId", "postId", "attachmentId")

read, err := client.GetThreadPostAttachment(ctx, id, threadpostattachment.DefaultGetThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostAttachmentClient.GetThreadPostAttachmentsCount`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

read, err := client.GetThreadPostAttachmentsCount(ctx, id, threadpostattachment.DefaultGetThreadPostAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostAttachmentClient.ListThreadPostAttachments`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostID("groupId", "conversationThreadId", "postId")

// alternatively `client.ListThreadPostAttachments(ctx, id, threadpostattachment.DefaultListThreadPostAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListThreadPostAttachmentsComplete(ctx, id, threadpostattachment.DefaultListThreadPostAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
