
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfoldermessageattachment` Documentation

The `usermailfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfoldermessageattachment"
```


### Client Initialization

```go
client := usermailfoldermessageattachment.NewUserMailFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderMessageAttachmentClient.CreateUserByIdMailFolderByIdMessageByIdAttachment`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageAttachmentClient.CreateUserByIdMailFolderByIdMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessageattachment.CreateUserByIdMailFolderByIdMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageAttachmentClient.DeleteUserByIdMailFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageAttachmentID("userIdValue", "mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdMailFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageAttachmentClient.GetUserByIdMailFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageAttachmentID("userIdValue", "mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.GetUserByIdMailFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageAttachmentClient.GetUserByIdMailFolderByIdMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageAttachmentClient.ListUserByIdMailFolderByIdMessageByIdAttachments`

```go
ctx := context.TODO()
id := usermailfoldermessageattachment.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

// alternatively `client.ListUserByIdMailFolderByIdMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFolderByIdMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
