
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplytoattachment` Documentation

The `conversationthreadpostinreplytoattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplytoattachment"
```


### Client Initialization

```go
client := conversationthreadpostinreplytoattachment.NewConversationThreadPostInReplyToAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.CreateConversationThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostinreplytoattachment.Attachment{
	// ...
}


read, err := client.CreateConversationThreadPostInReplyToAttachment(ctx, id, payload, conversationthreadpostinreplytoattachment.DefaultCreateConversationThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.CreateConversationThreadPostInReplyToAttachmentsUploadSession`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

payload := conversationthreadpostinreplytoattachment.CreateConversationThreadPostInReplyToAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateConversationThreadPostInReplyToAttachmentsUploadSession(ctx, id, payload, conversationthreadpostinreplytoattachment.DefaultCreateConversationThreadPostInReplyToAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.DeleteConversationThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId")

read, err := client.DeleteConversationThreadPostInReplyToAttachment(ctx, id, conversationthreadpostinreplytoattachment.DefaultDeleteConversationThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.GetConversationThreadPostInReplyToAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostIdInReplyToAttachmentID("groupId", "conversationId", "conversationThreadId", "postId", "attachmentId")

read, err := client.GetConversationThreadPostInReplyToAttachment(ctx, id, conversationthreadpostinreplytoattachment.DefaultGetConversationThreadPostInReplyToAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.GetConversationThreadPostInReplyToAttachmentsCount`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

read, err := client.GetConversationThreadPostInReplyToAttachmentsCount(ctx, id, conversationthreadpostinreplytoattachment.DefaultGetConversationThreadPostInReplyToAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostInReplyToAttachmentClient.ListConversationThreadPostInReplyToAttachments`

```go
ctx := context.TODO()
id := conversationthreadpostinreplytoattachment.NewGroupIdConversationIdThreadIdPostID("groupId", "conversationId", "conversationThreadId", "postId")

// alternatively `client.ListConversationThreadPostInReplyToAttachments(ctx, id, conversationthreadpostinreplytoattachment.DefaultListConversationThreadPostInReplyToAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListConversationThreadPostInReplyToAttachmentsComplete(ctx, id, conversationthreadpostinreplytoattachment.DefaultListConversationThreadPostInReplyToAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
