
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessageattachment` Documentation

The `mailfolderchildfoldermessageattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessageattachment"
```


### Client Initialization

```go
client := mailfolderchildfoldermessageattachment.NewMailFolderChildFolderMessageAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.CreateMailFolderChildFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageAttachment(ctx, id, payload, mailfolderchildfoldermessageattachment.DefaultCreateMailFolderChildFolderMessageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.CreateMailFolderChildFolderMessageAttachmentsUploadSession`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessageattachment.CreateMailFolderChildFolderMessageAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageAttachmentsUploadSession(ctx, id, payload, mailfolderchildfoldermessageattachment.DefaultCreateMailFolderChildFolderMessageAttachmentsUploadSessionOperationOptions())
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
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageIdAttachmentID("mailFolderId", "mailFolderId1", "messageId", "attachmentId")

read, err := client.DeleteMailFolderChildFolderMessageAttachment(ctx, id, mailfolderchildfoldermessageattachment.DefaultDeleteMailFolderChildFolderMessageAttachmentOperationOptions())
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
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageIdAttachmentID("mailFolderId", "mailFolderId1", "messageId", "attachmentId")

read, err := client.GetMailFolderChildFolderMessageAttachment(ctx, id, mailfolderchildfoldermessageattachment.DefaultGetMailFolderChildFolderMessageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageAttachmentClient.GetMailFolderChildFolderMessageAttachmentsCount`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderId", "mailFolderId1", "messageId")

read, err := client.GetMailFolderChildFolderMessageAttachmentsCount(ctx, id, mailfolderchildfoldermessageattachment.DefaultGetMailFolderChildFolderMessageAttachmentsCountOperationOptions())
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
id := mailfolderchildfoldermessageattachment.NewMeMailFolderIdChildFolderIdMessageID("mailFolderId", "mailFolderId1", "messageId")

// alternatively `client.ListMailFolderChildFolderMessageAttachments(ctx, id, mailfolderchildfoldermessageattachment.DefaultListMailFolderChildFolderMessageAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderChildFolderMessageAttachmentsComplete(ctx, id, mailfolderchildfoldermessageattachment.DefaultListMailFolderChildFolderMessageAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
