
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstanceattachment.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultDeleteCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, calendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachmentsCount(ctx, id, calendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultGetCalendarViewExceptionOccurrenceInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id, calendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id, calendarcalendarviewexceptionoccurrenceinstanceattachment.DefaultListCalendarViewExceptionOccurrenceInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
