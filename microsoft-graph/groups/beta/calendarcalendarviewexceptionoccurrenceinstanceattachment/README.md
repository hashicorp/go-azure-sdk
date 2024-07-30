
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `calendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstanceattachment.CreateCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarCalendarViewExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarCalendarViewExceptionOccurrenceInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListCalendarCalendarViewExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarCalendarViewExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
