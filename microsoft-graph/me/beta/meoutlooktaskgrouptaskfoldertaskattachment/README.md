
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskgrouptaskfoldertaskattachment` Documentation

The `meoutlooktaskgrouptaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskgrouptaskfoldertaskattachment"
```


### Client Initialization

```go
client := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.CreateMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := meoutlooktaskgrouptaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.CreateMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := meoutlooktaskgrouptaskfoldertaskattachment.CreateMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.DeleteMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.GetMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.GetMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskGroupTaskFolderTaskAttachmentClient.ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOutlookTaskGroupByIdTaskFolderByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
