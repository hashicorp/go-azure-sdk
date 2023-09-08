
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceexceptionoccurrenceattachment` Documentation

The `meeventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrenceattachment.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.DeleteMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.GetMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.GetMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceAttachmentClient.ListMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
