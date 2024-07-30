
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoattachment` Documentation

The `threadpostinreplytoattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoattachment"
```


### Client Initialization

```go
client := threadpostinreplytoattachment.NewThreadPostInReplyToAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.CreateThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplytoattachment.Attachment{
	// ...
}


read, err := client.CreateThreadPostInReplyToAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.CreateThreadPostInReplyToAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

payload := threadpostinreplytoattachment.CreateThreadPostInReplyToAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateThreadPostInReplyToAttachmentCreateUploadSession(ctx, id, payload)
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
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostIdInReplyToAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteThreadPostInReplyToAttachment(ctx, id)
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
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostIdInReplyToAttachmentID("groupIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetThreadPostInReplyToAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ThreadPostInReplyToAttachmentClient.GetThreadPostInReplyToAttachmentCount`

```go
ctx := context.TODO()
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetThreadPostInReplyToAttachmentCount(ctx, id)
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
id := threadpostinreplytoattachment.NewGroupIdThreadIdPostID("groupIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListThreadPostInReplyToAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListThreadPostInReplyToAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
