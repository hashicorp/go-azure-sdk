
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessageattachment` Documentation

The `mailfoldermessageattachment` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessageattachment"
```


### Client Initialization

```go
client := mailfoldermessageattachment.NewMailFolderMessageAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderMessageAttachmentClient.CreateMailFolderMessageAttachment`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMailFolderMessageAttachment(ctx, id, payload, mailfoldermessageattachment.DefaultCreateMailFolderMessageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.CreateMailFolderMessageAttachmentsUploadSession`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessageattachment.CreateMailFolderMessageAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateMailFolderMessageAttachmentsUploadSession(ctx, id, payload, mailfoldermessageattachment.DefaultCreateMailFolderMessageAttachmentsUploadSessionOperationOptions())
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
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageIdAttachmentID("userId", "mailFolderId", "messageId", "attachmentId")

read, err := client.DeleteMailFolderMessageAttachment(ctx, id, mailfoldermessageattachment.DefaultDeleteMailFolderMessageAttachmentOperationOptions())
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
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageIdAttachmentID("userId", "mailFolderId", "messageId", "attachmentId")

read, err := client.GetMailFolderMessageAttachment(ctx, id, mailfoldermessageattachment.DefaultGetMailFolderMessageAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageAttachmentClient.GetMailFolderMessageAttachmentsCount`

```go
ctx := context.TODO()
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

read, err := client.GetMailFolderMessageAttachmentsCount(ctx, id, mailfoldermessageattachment.DefaultGetMailFolderMessageAttachmentsCountOperationOptions())
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
id := mailfoldermessageattachment.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

// alternatively `client.ListMailFolderMessageAttachments(ctx, id, mailfoldermessageattachment.DefaultListMailFolderMessageAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderMessageAttachmentsComplete(ctx, id, mailfoldermessageattachment.DefaultListMailFolderMessageAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
