
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrenceattachment` Documentation

The `calendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendareventinstanceexceptionoccurrenceattachment.NewCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, payload)
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
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrenceattachment.CreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
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
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

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
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

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
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

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
id := calendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
