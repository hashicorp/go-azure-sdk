
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolderchildfoldermessageattachment` Documentation

The `memailfolderchildfoldermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolderchildfoldermessageattachment"
```


### Client Initialization

```go
client := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdAttachment`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessageattachment.Attachment{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessageattachment.CreateMeMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.DeleteMeMailFolderByIdChildFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageAttachmentID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteMeMailFolderByIdChildFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.GetMeMailFolderByIdChildFolderByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageAttachmentID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "attachmentIdValue")

read, err := client.GetMeMailFolderByIdChildFolderByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.GetMeMailFolderByIdChildFolderByIdMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMeMailFolderByIdChildFolderByIdMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageAttachmentClient.ListMeMailFolderByIdChildFolderByIdMessageByIdAttachments`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

// alternatively `client.ListMeMailFolderByIdChildFolderByIdMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMailFolderByIdChildFolderByIdMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
