
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `calendargroupcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.CreateCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendargroupcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarGroupCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
