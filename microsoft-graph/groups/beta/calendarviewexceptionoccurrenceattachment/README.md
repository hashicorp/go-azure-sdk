
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceattachment` Documentation

The `calendarviewexceptionoccurrenceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceattachment.NewCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachment(ctx, id, payload, calendarviewexceptionoccurrenceattachment.DefaultCreateCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarviewexceptionoccurrenceattachment.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload, calendarviewexceptionoccurrenceattachment.DefaultCreateCalendarViewExceptionOccurrenceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.DeleteCalendarViewExceptionOccurrenceAttachment(ctx, id, calendarviewexceptionoccurrenceattachment.DefaultDeleteCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.GetCalendarViewExceptionOccurrenceAttachment(ctx, id, calendarviewexceptionoccurrenceattachment.DefaultGetCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrenceAttachmentsCount(ctx, id, calendarviewexceptionoccurrenceattachment.DefaultGetCalendarViewExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.ListCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewExceptionOccurrenceAttachments(ctx, id, calendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id, calendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
