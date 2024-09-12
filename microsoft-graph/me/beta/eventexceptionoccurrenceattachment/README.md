
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


### Example Usage: `EventExceptionOccurrenceAttachmentClient.CreateEventExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := eventexceptionoccurrenceattachment.CreateEventExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
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

read, err := client.DeleteEventExceptionOccurrenceAttachment(ctx, id, eventexceptionoccurrenceattachment.DefaultDeleteEventExceptionOccurrenceAttachmentOperationOptions())
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

read, err := client.GetEventExceptionOccurrenceAttachment(ctx, id, eventexceptionoccurrenceattachment.DefaultGetEventExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.GetEventExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewMeEventIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetEventExceptionOccurrenceAttachmentsCount(ctx, id, eventexceptionoccurrenceattachment.DefaultGetEventExceptionOccurrenceAttachmentsCountOperationOptions())
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

// alternatively `client.ListEventExceptionOccurrenceAttachments(ctx, id, eventexceptionoccurrenceattachment.DefaultListEventExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceAttachmentsComplete(ctx, id, eventexceptionoccurrenceattachment.DefaultListEventExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
