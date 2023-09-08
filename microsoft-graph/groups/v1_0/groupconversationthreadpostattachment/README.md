
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupconversationthreadpostattachment` Documentation

The `groupconversationthreadpostattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupconversationthreadpostattachment"
```


### Client Initialization

```go
client := groupconversationthreadpostattachment.NewGroupConversationThreadPostAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.CreateGroupByIdConversationByIdThreadByIdPostByIdAttachment`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.CreateGroupByIdConversationByIdThreadByIdPostByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := groupconversationthreadpostattachment.CreateGroupByIdConversationByIdThreadByIdPostByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdConversationByIdThreadByIdPostByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.DeleteGroupByIdConversationByIdThreadByIdPostByIdAttachmentById`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdConversationByIdThreadByIdPostByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.GetGroupByIdConversationByIdThreadByIdPostByIdAttachmentById`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.GetGroupByIdConversationByIdThreadByIdPostByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetGroupByIdConversationByIdThreadByIdPostByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupConversationThreadPostAttachmentClient.ListGroupByIdConversationByIdThreadByIdPostByIdAttachments`

```go
ctx := context.TODO()
id := groupconversationthreadpostattachment.NewGroupConversationThreadPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListGroupByIdConversationByIdThreadByIdPostByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdConversationByIdThreadByIdPostByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
