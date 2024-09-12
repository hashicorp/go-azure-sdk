
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstanceattachment` Documentation

The `calendargroupcalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstanceattachment.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendareventexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, calendargroupcalendareventexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendargroupcalendareventexceptionoccurrenceinstanceattachment.DefaultGetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachments(ctx, id, calendargroupcalendareventexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendargroupcalendareventexceptionoccurrenceinstanceattachment.DefaultListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
