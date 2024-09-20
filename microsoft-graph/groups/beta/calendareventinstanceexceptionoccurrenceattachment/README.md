
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceattachment` Documentation

The `calendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendareventinstanceexceptionoccurrenceattachment.NewCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, payload, calendareventinstanceexceptionoccurrenceattachment.DefaultCreateCalendarEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventinstanceexceptionoccurrenceattachment.CreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload, calendareventinstanceexceptionoccurrenceattachment.DefaultCreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.DeleteCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.GetCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultGetCalendarEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarEventInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarEventInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultGetCalendarEventInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.ListCalendarEventInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
