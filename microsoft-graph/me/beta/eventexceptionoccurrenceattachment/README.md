
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceattachment` Documentation

The `eventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := eventexceptionoccurrenceattachment.NewEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.CreateEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := eventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.CreateEventExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := eventexceptionoccurrenceattachment.CreateEventExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.DeleteEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceIdAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteEventExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.GetEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceIdAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetEventExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.GetEventExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetEventExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.ListEventExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListEventExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
