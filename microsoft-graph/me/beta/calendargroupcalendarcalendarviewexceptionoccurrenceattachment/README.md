
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrenceattachment` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
