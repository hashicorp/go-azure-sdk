
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostattachment` Documentation

The `threadpostattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostattachment"
```


### Client Initialization

```go
client := threadpostattachment.NewThreadPostAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostAttachmentClient.CreateThreadPostAttachment`

```go
ctx := context.TODO()
id := threadpostattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostattachment.Attachment{
	// ...
}


read, err := client.CreateThreadPostAttachment(ctx, id, payload)
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
id := threadpostattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostattachment.CreateThreadPostAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateThreadPostAttachmentsUploadSession(ctx, id, payload)
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
id := threadpostattachment.NewGroupIdThreadIdPostIdAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

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
id := threadpostattachment.NewGroupIdThreadIdPostIdAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

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
id := threadpostattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

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
id := threadpostattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListThreadPostAttachments(ctx, id, threadpostattachment.DefaultListThreadPostAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListThreadPostAttachmentsComplete(ctx, id, threadpostattachment.DefaultListThreadPostAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
