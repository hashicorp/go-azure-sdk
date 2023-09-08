
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meeventattachment` Documentation

The `meeventattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meeventattachment"
```


### Client Initialization

```go
client := meeventattachment.NewMeEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventAttachmentClient.CreateMeEventByIdAttachment`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventID("eventIdValue")

payload := meeventattachment.Attachment{
	// ...
}


read, err := client.CreateMeEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventAttachmentClient.CreateMeEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventID("eventIdValue")

payload := meeventattachment.CreateMeEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventAttachmentClient.DeleteMeEventByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventAttachmentClient.GetMeEventByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventAttachmentClient.GetMeEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventID("eventIdValue")

read, err := client.GetMeEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventAttachmentClient.ListMeEventByIdAttachments`

```go
ctx := context.TODO()
id := meeventattachment.NewMeEventID("eventIdValue")

// alternatively `client.ListMeEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
