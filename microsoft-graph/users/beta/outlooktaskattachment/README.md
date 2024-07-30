
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskattachment` Documentation

The `outlooktaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskattachment"
```


### Client Initialization

```go
client := outlooktaskattachment.NewOutlookTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutlookTaskAttachmentClient.CreateOutlookTaskAttachment`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewUserIdOutlookTaskID("userIdValue", "outlookTaskIdValue")

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


### Example Usage: `OutlookTaskAttachmentClient.CreateOutlookTaskAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewUserIdOutlookTaskID("userIdValue", "outlookTaskIdValue")

payload := outlooktaskattachment.CreateOutlookTaskAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateOutlookTaskAttachmentCreateUploadSession(ctx, id, payload)
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
id := outlooktaskattachment.NewUserIdOutlookTaskIdAttachmentID("userIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteOutlookTaskAttachment(ctx, id)
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
id := outlooktaskattachment.NewUserIdOutlookTaskIdAttachmentID("userIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetOutlookTaskAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OutlookTaskAttachmentClient.GetOutlookTaskAttachmentCount`

```go
ctx := context.TODO()
id := outlooktaskattachment.NewUserIdOutlookTaskID("userIdValue", "outlookTaskIdValue")

read, err := client.GetOutlookTaskAttachmentCount(ctx, id)
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
id := outlooktaskattachment.NewUserIdOutlookTaskID("userIdValue", "outlookTaskIdValue")

// alternatively `client.ListOutlookTaskAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListOutlookTaskAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
