
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceattachment` Documentation

The `calendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendareventexceptionoccurrenceinstanceattachment.NewCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, payload, calendareventexceptionoccurrenceinstanceattachment.DefaultCreateCalendarEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendareventexceptionoccurrenceinstanceattachment.CreateCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload, calendareventexceptionoccurrenceinstanceattachment.DefaultCreateCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.DeleteCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, calendareventexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.GetCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, calendareventexceptionoccurrenceinstanceattachment.DefaultGetCalendarEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarEventExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarEventExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendareventexceptionoccurrenceinstanceattachment.DefaultGetCalendarEventExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.ListCalendarEventExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

// alternatively `client.ListCalendarEventExceptionOccurrenceInstanceAttachments(ctx, id, calendareventexceptionoccurrenceinstanceattachment.DefaultListCalendarEventExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendareventexceptionoccurrenceinstanceattachment.DefaultListCalendarEventExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
