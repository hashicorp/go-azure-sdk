
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceinstanceattachment` Documentation

The `eventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := eventexceptionoccurrenceinstanceattachment.NewEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.CreateEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.CreateEventExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := eventexceptionoccurrenceinstanceattachment.CreateEventExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
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
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteEventExceptionOccurrenceInstanceAttachment(ctx, id)
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
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceIdAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetEventExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceAttachmentClient.GetEventExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetEventExceptionOccurrenceInstanceAttachmentCount(ctx, id)
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
id := eventexceptionoccurrenceinstanceattachment.NewMeEventIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListEventExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
