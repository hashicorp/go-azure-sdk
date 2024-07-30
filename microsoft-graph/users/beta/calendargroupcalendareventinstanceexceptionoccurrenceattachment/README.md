
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrenceattachment` Documentation

The `calendargroupcalendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

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


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrenceattachment.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
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
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id)
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
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentCount(ctx, id)
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
id := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
