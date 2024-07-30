
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessageattachment` Documentation

The `mailfolderchildfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessageattachment"
```


### Client Initialization

```go
client := mailfolderchildfoldermessageattachment.NewMailFolderChildFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.CreateMailFolderChildFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.CreateMailFolderChildFolderMessageAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessageattachment.CreateMailFolderChildFolderMessageAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.DeleteMailFolderChildFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageIdAttachmentID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteMailFolderChildFolderMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.GetMailFolderChildFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageIdAttachmentID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.GetMailFolderChildFolderMessageAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.GetMailFolderChildFolderMessageAttachmentCount`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMailFolderChildFolderMessageAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.ListMailFolderChildFolderMessageAttachments`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

// alternatively `client.ListMailFolderChildFolderMessageAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMailFolderChildFolderMessageAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
