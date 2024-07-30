
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventexceptionoccurrenceinstanceattachment` Documentation

The `calendargroupcalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

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


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstanceattachment.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
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
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id)
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
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentCount(ctx, id)
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
id := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
