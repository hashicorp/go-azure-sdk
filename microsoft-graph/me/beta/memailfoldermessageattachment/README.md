
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfoldermessageattachment` Documentation

The `memailfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfoldermessageattachment"
```


### Client Initialization

```go
client := memailfoldermessageattachment.NewMeMailFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderMessageAttachmentClient.CreateMeMailFolderByIdMessageByIdAttachment`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageAttachmentClient.CreateMeMailFolderByIdMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessageattachment.CreateMeMailFolderByIdMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageAttachmentClient.DeleteMeMailFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageAttachmentID("mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteMeMailFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageAttachmentClient.GetMeMailFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageAttachmentID("mailFolderIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.GetMeMailFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageAttachmentClient.GetMeMailFolderByIdMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMeMailFolderByIdMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageAttachmentClient.ListMeMailFolderByIdMessageByIdAttachments`

```go
ctx := context.TODO()
id := memailfoldermessageattachment.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

// alternatively `client.ListMeMailFolderByIdMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMailFolderByIdMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
