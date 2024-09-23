
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceattachment` Documentation

The `calendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceattachment.NewCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachment(ctx, id, payload, calendarcalendarviewexceptionoccurrenceattachment.DefaultCreateCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

payload := calendarcalendarviewexceptionoccurrenceattachment.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload, calendarcalendarviewexceptionoccurrenceattachment.DefaultCreateCalendarViewExceptionOccurrenceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.DeleteCalendarViewExceptionOccurrenceAttachment(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultDeleteCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.GetCalendarViewExceptionOccurrenceAttachment(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewExceptionOccurrenceAttachmentsCount(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarViewExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.ListCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewExceptionOccurrenceAttachments(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
