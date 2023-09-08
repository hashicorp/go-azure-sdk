
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpostinreplytoattachment` Documentation

The `groupconversationthreadpostinreplytoattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupconversationthreadpostinreplytoattachment"
```


### Client Initialization

```go
client := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostInReplyToAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachment`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostinreplytoattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostinreplytoattachment.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.DeleteGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentById`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostInReplyToAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentById`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostInReplyToAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentCount`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostInReplyToAttachmentClient.ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachments`

```go
ctx := context.TODO()
id := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdConversationByIdThreadByIdPostByIdInReplyToAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
