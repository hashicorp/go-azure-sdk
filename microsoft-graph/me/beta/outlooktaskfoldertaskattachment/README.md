
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskfoldertaskattachment` Documentation

The `outlooktaskfoldertaskattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskfoldertaskattachment"
```


### Client Initialization

```go
client := outlooktaskfoldertaskattachment.NewOutlookTaskFolderTaskAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.CreateOutlookTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId")

payload := outlooktaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateOutlookTaskFolderTaskAttachment(ctx, id, payload, outlooktaskfoldertaskattachment.DefaultCreateOutlookTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.CreateOutlookTaskFolderTaskAttachmentsUploadSession`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId")

payload := outlooktaskfoldertaskattachment.CreateOutlookTaskFolderTaskAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskFolderTaskAttachmentsUploadSession(ctx, id, payload, outlooktaskfoldertaskattachment.DefaultCreateOutlookTaskFolderTaskAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.DeleteOutlookTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderId", "outlookTaskId", "attachmentId")

read, err := client.DeleteOutlookTaskFolderTaskAttachment(ctx, id, outlooktaskfoldertaskattachment.DefaultDeleteOutlookTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.GetOutlookTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderId", "outlookTaskId", "attachmentId")

read, err := client.GetOutlookTaskFolderTaskAttachment(ctx, id, outlooktaskfoldertaskattachment.DefaultGetOutlookTaskFolderTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.GetOutlookTaskFolderTaskAttachmentsCount`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId")

read, err := client.GetOutlookTaskFolderTaskAttachmentsCount(ctx, id, outlooktaskfoldertaskattachment.DefaultGetOutlookTaskFolderTaskAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.ListOutlookTaskFolderTaskAttachments`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderId", "outlookTaskId")

// alternatively `client.ListOutlookTaskFolderTaskAttachments(ctx, id, outlooktaskfoldertaskattachment.DefaultListOutlookTaskFolderTaskAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListOutlookTaskFolderTaskAttachmentsComplete(ctx, id, outlooktaskfoldertaskattachment.DefaultListOutlookTaskFolderTaskAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
