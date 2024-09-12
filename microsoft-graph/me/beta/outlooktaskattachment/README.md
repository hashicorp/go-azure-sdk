
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskattachment` Documentation

The `outlooktaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/outlooktaskattachment"
```


### Client Initialization

```go
client := outlooktaskattachment.NewOutlookTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskAttachmentClient.CreateOutlookTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

payload := outlooktaskattachment.Attachment{
	// ...
}


read, err := client.CreateOutlookTaskAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.CreateOutlookTaskAttachmentsUploadSession`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

payload := outlooktaskattachment.CreateOutlookTaskAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.DeleteOutlookTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskIdAttachmentID("outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteOutlookTaskAttachment(ctx, id, outlooktaskattachment.DefaultDeleteOutlookTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.GetOutlookTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskIdAttachmentID("outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetOutlookTaskAttachment(ctx, id, outlooktaskattachment.DefaultGetOutlookTaskAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.GetOutlookTaskAttachmentsCount`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

read, err := client.GetOutlookTaskAttachmentsCount(ctx, id, outlooktaskattachment.DefaultGetOutlookTaskAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.ListOutlookTaskAttachments`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

// alternatively `client.ListOutlookTaskAttachments(ctx, id, outlooktaskattachment.DefaultListOutlookTaskAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListOutlookTaskAttachmentsComplete(ctx, id, outlooktaskattachment.DefaultListOutlookTaskAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
