
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceinstanceattachment.NewCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstanceattachment.CreateCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarViewExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
