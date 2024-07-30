
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mailfoldermessageattachment` Documentation

The `mailfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mailfoldermessageattachment"
```


### Client Initialization

```go
client := mailfoldermessageattachment.NewMailFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderMessageAttachmentClient.CreateMailFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMailFolderMessageAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.CreateMailFolderMessageAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessageattachment.CreateMailFolderMessageAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMailFolderMessageAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.DeleteMailFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageIdAttachmentID("mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteMailFolderMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.GetMailFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageIdAttachmentID("mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.GetMailFolderMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.GetMailFolderMessageAttachmentCount`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMailFolderMessageAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.ListMailFolderMessageAttachments`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

// alternatively `client.ListMailFolderMessageAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMailFolderMessageAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
