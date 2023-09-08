
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceattachment` Documentation

The `meeventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceattachment"
```


### Client Initialization

```go
client := meeventinstanceattachment.NewMeEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventInstanceAttachmentClient.CreateMeEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceAttachmentClient.CreateMeEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstanceattachment.CreateMeEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceAttachmentClient.DeleteMeEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceAttachmentClient.GetMeEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceAttachmentClient.GetMeEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceAttachmentClient.ListMeEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := meeventinstanceattachment.NewMeEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
