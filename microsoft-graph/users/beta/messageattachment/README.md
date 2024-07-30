
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/messageattachment` Documentation

The `messageattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/messageattachment"
```


### Client Initialization

```go
client := messageattachment.NewMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageAttachmentClient.CreateMessageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := messageattachment.Attachment{
	// ...
}


read, err := client.CreateMessageAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.CreateMessageAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := messageattachment.CreateMessageAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMessageAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.DeleteMessageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageIdAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.GetMessageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageIdAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.GetMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.GetMessageAttachmentCount`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageID("userIdValue", "messageIdValue")

read, err := client.GetMessageAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.ListMessageAttachments`

```go
ctx := context.TODO()
id := messageattachment.NewUserIdMessageID("userIdValue", "messageIdValue")

// alternatively `client.ListMessageAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMessageAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
