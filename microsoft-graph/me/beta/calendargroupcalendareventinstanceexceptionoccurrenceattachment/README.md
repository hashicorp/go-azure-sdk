
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventinstanceexceptionoccurrenceattachment` Documentation

The `calendargroupcalendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrenceattachment.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendareventinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendareventinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendargroupcalendareventinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id, calendargroupcalendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendargroupcalendareventinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
