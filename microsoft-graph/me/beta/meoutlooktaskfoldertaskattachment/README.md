
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskfoldertaskattachment` Documentation

The `meoutlooktaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskfoldertaskattachment"
```


### Client Initialization

```go
client := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.CreateMeOutlookTaskFolderByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := meoutlooktaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateMeOutlookTaskFolderByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.CreateMeOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := meoutlooktaskfoldertaskattachment.CreateMeOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeOutlookTaskFolderByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.DeleteMeOutlookTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskAttachmentID("outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteMeOutlookTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.GetMeOutlookTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskAttachmentID("outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetMeOutlookTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.GetMeOutlookTaskFolderByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetMeOutlookTaskFolderByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskFolderTaskAttachmentClient.ListMeOutlookTaskFolderByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListMeOutlookTaskFolderByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOutlookTaskFolderByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
