
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceattachment` Documentation

The `calendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceattachment.NewCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachment(ctx, id, payload)
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
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrenceattachment.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
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
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

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
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

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
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

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
id := calendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarCalendarViewIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceAttachments(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id, calendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarViewExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
