
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceinstanceattachment.NewCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload, calendarviewexceptionoccurrenceinstanceattachment.DefaultCreateCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewexceptionoccurrenceinstanceattachment.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload, calendarviewexceptionoccurrenceinstanceattachment.DefaultCreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.DeleteCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendarviewexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarViewExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id, calendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
