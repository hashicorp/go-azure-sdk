
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostattachment` Documentation

The `conversationthreadpostattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostattachment"
```


### Client Initialization

```go
client := conversationthreadpostattachment.NewConversationThreadPostAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConversationThreadPostAttachmentClient.CreateConversationThreadPostAttachment`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpostattachment.Attachment{
	// ...
}


read, err := client.CreateConversationThreadPostAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.CreateConversationThreadPostAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

payload := conversationthreadpostattachment.CreateConversationThreadPostAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateConversationThreadPostAttachmentCreateUploadSession(ctx, id, payload)
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
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostIdAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.DeleteConversationThreadPostAttachment(ctx, id)
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
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostIdAttachmentID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue", "attachmentIdValue")

read, err := client.GetConversationThreadPostAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConversationThreadPostAttachmentClient.GetConversationThreadPostAttachmentCount`

```go
ctx := context.TODO()
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

read, err := client.GetConversationThreadPostAttachmentCount(ctx, id)
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
id := conversationthreadpostattachment.NewGroupIdConversationIdThreadIdPostID("groupIdValue", "conversationIdValue", "conversationThreadIdValue", "postIdValue")

// alternatively `client.ListConversationThreadPostAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListConversationThreadPostAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
