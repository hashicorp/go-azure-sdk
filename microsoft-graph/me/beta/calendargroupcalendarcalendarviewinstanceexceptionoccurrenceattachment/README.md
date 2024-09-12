
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
