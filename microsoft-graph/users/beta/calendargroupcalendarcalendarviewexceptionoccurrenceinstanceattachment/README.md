
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


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
