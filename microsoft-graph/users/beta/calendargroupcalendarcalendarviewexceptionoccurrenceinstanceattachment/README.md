
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
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


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
