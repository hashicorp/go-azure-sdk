
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/messageattachment` Documentation

The `messageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/messageattachment"
```


### Client Initialization

```go
client := messageattachment.NewMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageAttachmentClient.CreateCreatessageAttachmentsUploadSession`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageID("messageIdValue")

payload := messageattachment.CreateCreatessageAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCreatessageAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.CreatessageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageID("messageIdValue")

payload := messageattachment.Attachment{
	// ...
}


read, err := client.CreatessageAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.DeletessageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageIdAttachmentID("messageIdValue", "attachmentIdValue")

read, err := client.DeletessageAttachment(ctx, id, messageattachment.DefaultDeletessageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.GetssageAttachment`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageIdAttachmentID("messageIdValue", "attachmentIdValue")

read, err := client.GetssageAttachment(ctx, id, messageattachment.DefaultGetssageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.GetssageAttachmentsCount`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageID("messageIdValue")

read, err := client.GetssageAttachmentsCount(ctx, id, messageattachment.DefaultGetssageAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageAttachmentClient.ListssageAttachments`

```go
ctx := context.TODO()
id := messageattachment.NewMeMessageID("messageIdValue")

// alternatively `client.ListssageAttachments(ctx, id, messageattachment.DefaultListssageAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListssageAttachmentsComplete(ctx, id, messageattachment.DefaultListssageAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
