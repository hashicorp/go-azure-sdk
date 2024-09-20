
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceinstanceattachment` Documentation

The `eventexceptionoccurrenceinstanceattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := eventexceptionoccurrenceinstanceattachment.NewEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.CreateEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceAttachment(ctx, id, payload, eventexceptionoccurrenceinstanceattachment.DefaultCreateEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.CreateEventExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstanceattachment.CreateEventExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload, eventexceptionoccurrenceinstanceattachment.DefaultCreateEventExceptionOccurrenceInstanceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.DeleteEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.DeleteEventExceptionOccurrenceInstanceAttachment(ctx, id, eventexceptionoccurrenceinstanceattachment.DefaultDeleteEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.GetEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.GetEventExceptionOccurrenceInstanceAttachment(ctx, id, eventexceptionoccurrenceinstanceattachment.DefaultGetEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.GetEventExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventId", "eventId1", "eventId2")

read, err := client.GetEventExceptionOccurrenceInstanceAttachmentsCount(ctx, id, eventexceptionoccurrenceinstanceattachment.DefaultGetEventExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.ListEventExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventId", "eventId1", "eventId2")

// alternatively `client.ListEventExceptionOccurrenceInstanceAttachments(ctx, id, eventexceptionoccurrenceinstanceattachment.DefaultListEventExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, eventexceptionoccurrenceinstanceattachment.DefaultListEventExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
