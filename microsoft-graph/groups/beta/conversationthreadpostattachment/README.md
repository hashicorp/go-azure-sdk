
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostattachment` Documentation

The `conversationthreadpostattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostattachment"
```


### Client Initialization

```go
client := conversationthreadpostattachment.NewConversationThreadPostAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostAttachmentClient.CreateConversationThreadPostAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostattachment.Attachment{
	// ...
}


read, err := client.CreateConversationThreadPostAttachment(ctx, id, payload, conversationthreadpostattachment.DefaultCreateConversationThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.CreateConversationThreadPostAttachmentsUploadSession`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostattachment.CreateConversationThreadPostAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateConversationThreadPostAttachmentsUploadSession(ctx, id, payload, conversationthreadpostattachment.DefaultCreateConversationThreadPostAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.DeleteConversationThreadPostAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostIdAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId")

read, err := client.DeleteConversationThreadPostAttachment(ctx, id, conversationthreadpostattachment.DefaultDeleteConversationThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.GetConversationThreadPostAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostIdAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId")

read, err := client.GetConversationThreadPostAttachment(ctx, id, conversationthreadpostattachment.DefaultGetConversationThreadPostAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.GetConversationThreadPostAttachmentsCount`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

read, err := client.GetConversationThreadPostAttachmentsCount(ctx, id, conversationthreadpostattachment.DefaultGetConversationThreadPostAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.ListConversationThreadPostAttachments`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

// alternatively `client.ListConversationThreadPostAttachments(ctx, id, conversationthreadpostattachment.DefaultListConversationThreadPostAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListConversationThreadPostAttachmentsComplete(ctx, id, conversationthreadpostattachment.DefaultListConversationThreadPostAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
