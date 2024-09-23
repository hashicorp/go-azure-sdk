
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrenceattachment` Documentation

The `calendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarviewinstanceexceptionoccurrenceattachment.NewCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, payload, calendarviewinstanceexceptionoccurrenceattachment.DefaultCreateCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

payload := calendarviewinstanceexceptionoccurrenceattachment.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload, calendarviewinstanceexceptionoccurrenceattachment.DefaultCreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.DeleteCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendarviewinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "eventId2", "attachmentId")

read, err := client.GetCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarViewInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetCalendarViewInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarViewInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupId", "eventId", "eventId1", "eventId2")

// alternatively `client.ListCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id, calendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
