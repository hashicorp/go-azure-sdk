
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `calendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrenceattachment.CreateCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarCalendarViewInstanceExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
