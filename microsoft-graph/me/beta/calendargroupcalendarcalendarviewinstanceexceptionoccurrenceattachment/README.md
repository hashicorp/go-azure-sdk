
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
