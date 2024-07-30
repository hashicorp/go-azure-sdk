
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceattachment` Documentation

The `calendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := calendareventexceptionoccurrenceinstanceattachment.NewCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstanceattachment.CreateCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarEventExceptionOccurrenceInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarEventExceptionOccurrenceInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.GetCalendarEventExceptionOccurrenceInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventExceptionOccurrenceInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceAttachmentClient.ListCalendarEventExceptionOccurrenceInstanceAttachments`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstanceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarEventExceptionOccurrenceInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrenceInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
