
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceattachment` Documentation

The `calendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendareventinstanceexceptionoccurrenceattachment.NewCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventinstanceexceptionoccurrenceattachment.CreateCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceExceptionOccurrenceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarEventInstanceExceptionOccurrenceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.GetCalendarEventInstanceExceptionOccurrenceAttachmentCount`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventInstanceExceptionOccurrenceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceExceptionOccurrenceAttachmentClient.ListCalendarEventInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendareventinstanceexceptionoccurrenceattachment.NewGroupIdCalendarEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarEventInstanceExceptionOccurrenceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
