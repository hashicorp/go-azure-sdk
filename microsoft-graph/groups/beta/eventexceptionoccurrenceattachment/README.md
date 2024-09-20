
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceattachment` Documentation

The `eventexceptionoccurrenceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := eventexceptionoccurrenceattachment.NewEventExceptionOccurrenceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceAttachmentClient.CreateEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := eventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAttachment(ctx, id, payload, eventexceptionoccurrenceattachment.DefaultCreateEventExceptionOccurrenceAttachmentOperationOptions())
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
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := eventexceptionoccurrenceattachment.CreateEventExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload, eventexceptionoccurrenceattachment.DefaultCreateEventExceptionOccurrenceAttachmentsUploadSessionOperationOptions())
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
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

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
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

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
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

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
id := eventexceptionoccurrenceattachment.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListEventExceptionOccurrenceAttachments(ctx, id, eventexceptionoccurrenceattachment.DefaultListEventExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceAttachmentsComplete(ctx, id, eventexceptionoccurrenceattachment.DefaultListEventExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
