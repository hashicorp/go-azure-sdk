
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceinstanceattachment` Documentation

The `meeventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventexceptionoccurrenceinstanceattachment.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.DeleteMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.GetMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.GetMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceInstanceAttachmentClient.ListMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
