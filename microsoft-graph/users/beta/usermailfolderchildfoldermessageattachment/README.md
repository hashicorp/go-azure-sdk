
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfolderchildfoldermessageattachment` Documentation

The `usermailfolderchildfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfolderchildfoldermessageattachment"
```


### Client Initialization

```go
client := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachment`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessageattachment.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.DeleteUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageAttachmentID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageAttachmentID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageAttachmentClient.ListUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachments`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

// alternatively `client.ListUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFolderByIdChildFolderByIdMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
