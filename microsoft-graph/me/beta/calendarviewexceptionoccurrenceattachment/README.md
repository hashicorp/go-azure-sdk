
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrenceattachment` Documentation

The `calendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceattachment.NewCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.CreateCalendarViewExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := calendarviewexceptionoccurrenceattachment.CreateCalendarViewExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.DeleteCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceIdAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceIdAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarViewExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.GetCalendarViewExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceAttachmentClient.ListCalendarViewExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceattachment.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
