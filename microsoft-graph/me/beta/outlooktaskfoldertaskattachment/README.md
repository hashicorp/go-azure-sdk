
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskfoldertaskattachment` Documentation

The `outlooktaskfoldertaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskfoldertaskattachment"
```


### Client Initialization

```go
client := outlooktaskfoldertaskattachment.NewOutlookTaskFolderTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.CreateOutlookTaskFolderTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := outlooktaskfoldertaskattachment.Attachment{
	// ...
}


read, err := client.CreateOutlookTaskFolderTaskAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.CreateOutlookTaskFolderTaskAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

payload := outlooktaskfoldertaskattachment.CreateOutlookTaskFolderTaskAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskFolderTaskAttachmentCreateUploadSession(ctx, id, payload)
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
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteOutlookTaskFolderTaskAttachment(ctx, id)
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
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskIdAttachmentID("outlookTaskFolderIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetOutlookTaskFolderTaskAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskFolderTaskAttachmentClient.GetOutlookTaskFolderTaskAttachmentCount`

```go
ctx := context.TODO()
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

read, err := client.GetOutlookTaskFolderTaskAttachmentCount(ctx, id)
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
id := outlooktaskfoldertaskattachment.NewMeOutlookTaskFolderIdTaskID("outlookTaskFolderIdValue", "outlookTaskIdValue")

// alternatively `client.ListOutlookTaskFolderTaskAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListOutlookTaskFolderTaskAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
