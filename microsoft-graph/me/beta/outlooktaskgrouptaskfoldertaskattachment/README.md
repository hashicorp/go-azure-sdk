
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskgrouptaskfoldertaskattachment` Documentation

The `outlooktaskgrouptaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskgrouptaskfoldertaskattachment"
```


### Client Initialization

```go
client := outlooktaskgrouptaskfoldertaskattachment.NewOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.CreateOutlookTaskGroupTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := outlooktaskgrouptaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateOutlookTaskGroupTaskFolderTaskAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.CreateOutlookTaskGroupTaskFolderTaskAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := outlooktaskgrouptaskfoldertaskattachment.CreateOutlookTaskGroupTaskFolderTaskAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskGroupTaskFolderTaskAttachmentCreateUploadSession(ctx, id, payload)
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
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteOutlookTaskGroupTaskFolderTaskAttachment(ctx, id)
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
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskIdAttachmentID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetOutlookTaskGroupTaskFolderTaskAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskGroupTaskFolderTaskAttachmentClient.GetOutlookTaskGroupTaskFolderTaskAttachmentCount`

```go
ctx := context.TODO()
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetOutlookTaskGroupTaskFolderTaskAttachmentCount(ctx, id)
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
id := outlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupIdTaskFolderIdTaskID("outlookTaskGroupIdValue", "outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListOutlookTaskGroupTaskFolderTaskAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListOutlookTaskGroupTaskFolderTaskAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
