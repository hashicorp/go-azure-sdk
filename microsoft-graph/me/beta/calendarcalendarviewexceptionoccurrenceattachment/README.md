
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceattachment` Documentation

The `calendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceattachment.NewCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarIdCalendarViewIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewexceptionoccurrenceattachment.CreateCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.GetCalendarCalendarViewExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceAttachmentClient.ListCalendarCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarCalendarViewExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
