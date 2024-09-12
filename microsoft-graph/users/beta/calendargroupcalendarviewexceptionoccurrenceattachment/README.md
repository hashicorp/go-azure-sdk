
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewexceptionoccurrenceattachment` Documentation

The `calendargroupcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarviewexceptionoccurrenceattachment.NewCalendarGroupCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewexceptionoccurrenceattachment.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarviewexceptionoccurrenceattachment.DefaultDeleteCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCount(ctx, id, calendargroupcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceAttachments(ctx, id, calendargroupcalendarviewexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id, calendargroupcalendarviewexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
