
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceattachment` Documentation

The `meeventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.CreateMeEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.CreateMeEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrenceattachment.CreateMeEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.DeleteMeEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.GetMeEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.GetMeEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceAttachmentClient.ListMeEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
