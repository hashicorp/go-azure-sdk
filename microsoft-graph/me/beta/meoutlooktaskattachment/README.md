
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskattachment` Documentation

The `meoutlooktaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meoutlooktaskattachment"
```


### Client Initialization

```go
client := meoutlooktaskattachment.NewMeOutlookTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOutlookTaskAttachmentClient.CreateMeOutlookTaskByIdAttachment`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

payload := meoutlooktaskattachment.Attachment{
	// ...
}


read, err := client.CreateMeOutlookTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskAttachmentClient.CreateMeOutlookTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

payload := meoutlooktaskattachment.CreateMeOutlookTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeOutlookTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskAttachmentClient.DeleteMeOutlookTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskAttachmentID("outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteMeOutlookTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskAttachmentClient.GetMeOutlookTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskAttachmentID("outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetMeOutlookTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskAttachmentClient.GetMeOutlookTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

read, err := client.GetMeOutlookTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOutlookTaskAttachmentClient.ListMeOutlookTaskByIdAttachments`

```go
ctx := context.TODO()
id := meoutlooktaskattachment.NewMeOutlookTaskID("outlookTaskIdValue")

// alternatively `client.ListMeOutlookTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOutlookTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
