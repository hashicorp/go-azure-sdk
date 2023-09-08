
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskgrouptaskfoldertaskattachment` Documentation

The `useroutlooktaskgrouptaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskgrouptaskfoldertaskattachment"
```


### Client Initialization

```go
client := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.CreateUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := useroutlooktaskgrouptaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.CreateUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := useroutlooktaskgrouptaskfoldertaskattachment.CreateUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.DeleteUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskAttachmentID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.GetUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskAttachmentID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.GetUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskGroupTaskFolderTaskAttachmentClient.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskID("userIdValue", "outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
