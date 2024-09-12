
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceattachment` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.Attachment{
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


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.CreateCalendarGroupCalendarViewExceptionOccurrenceAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceattachment.DefaultDeleteCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCount(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceAttachments(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
