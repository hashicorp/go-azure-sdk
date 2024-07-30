
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceattachment` Documentation

The `calendargroupcalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrenceattachment.NewCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateCalendarGroupCalendarEventExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventexceptionoccurrenceattachment.CreateCalendarGroupCalendarEventExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.DeleteCalendarGroupCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetCalendarGroupCalendarEventExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.ListCalendarGroupCalendarEventExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
