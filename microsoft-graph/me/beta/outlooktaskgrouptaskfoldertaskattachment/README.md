
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskgrouptaskfoldertaskattachment` Documentation

The `outlooktaskgrouptaskfoldertaskattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskgrouptaskfoldertaskattachment"
```


### Client Initialization

```go
client := outlooktaskgrouptaskfoldertaskattachment.NewOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.CreateOutlookTaskGroupTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId")

payload := outlooktaskgrouptaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateOutlookTaskGroupTaskFolderTaskAttachment(ctx, id, payload, outlooktaskgrouptaskfoldertaskattachment.DefaultCreateOutlookTaskGroupTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.CreateOutlookTaskGroupTaskFolderTaskAttachmentsUploadSession`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId")

payload := outlooktaskgrouptaskfoldertaskattachment.CreateOutlookTaskGroupTaskFolderTaskAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskGroupTaskFolderTaskAttachmentsUploadSession(ctx, id, payload, outlooktaskgrouptaskfoldertaskattachment.DefaultCreateOutlookTaskGroupTaskFolderTaskAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.DeleteOutlookTaskGroupTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId", "attachmentId")

read, err := client.DeleteOutlookTaskGroupTaskFolderTaskAttachment(ctx, id, outlooktaskgrouptaskfoldertaskattachment.DefaultDeleteOutlookTaskGroupTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.GetOutlookTaskGroupTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId", "attachmentId")

read, err := client.GetOutlookTaskGroupTaskFolderTaskAttachment(ctx, id, outlooktaskgrouptaskfoldertaskattachment.DefaultGetOutlookTaskGroupTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.GetOutlookTaskGroupTaskFolderTaskAttachmentsCount`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId")

read, err := client.GetOutlookTaskGroupTaskFolderTaskAttachmentsCount(ctx, id, outlooktaskgrouptaskfoldertaskattachment.DefaultGetOutlookTaskGroupTaskFolderTaskAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.ListOutlookTaskGroupTaskFolderTaskAttachments`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupId", "outlookTaskFolderId", "outlookTaskId")

// alternatively `client.ListOutlookTaskGroupTaskFolderTaskAttachments(ctx, id, outlooktaskgrouptaskfoldertaskattachment.DefaultListOutlookTaskGroupTaskFolderTaskAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListOutlookTaskGroupTaskFolderTaskAttachmentsComplete(ctx, id, outlooktaskgrouptaskfoldertaskattachment.DefaultListOutlookTaskGroupTaskFolderTaskAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
