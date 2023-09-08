
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskfoldertaskattachment` Documentation

The `useroutlooktaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskfoldertaskattachment"
```


### Client Initialization

```go
client := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.CreateUserByIdOutlookTaskFolderByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := useroutlooktaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskFolderByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.CreateUserByIdOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := useroutlooktaskfoldertaskattachment.CreateUserByIdOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.DeleteUserByIdOutlookTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskAttachmentID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdOutlookTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.GetUserByIdOutlookTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskAttachmentID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetUserByIdOutlookTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.GetUserByIdOutlookTaskFolderByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetUserByIdOutlookTaskFolderByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskFolderTaskAttachmentClient.ListUserByIdOutlookTaskFolderByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskID("userIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListUserByIdOutlookTaskFolderByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOutlookTaskFolderByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
