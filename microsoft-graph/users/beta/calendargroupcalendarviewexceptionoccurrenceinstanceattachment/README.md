
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendargroupcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
