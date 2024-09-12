
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceattachment` Documentation

The `calendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendareventexceptionoccurrenceattachment.NewCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.CreateCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.CreateCalendarEventExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventexceptionoccurrenceattachment.CreateCalendarEventExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.DeleteCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarEventExceptionOccurrenceAttachment(ctx, id, calendareventexceptionoccurrenceattachment.DefaultDeleteCalendarEventExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.GetCalendarEventExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarEventExceptionOccurrenceAttachment(ctx, id, calendareventexceptionoccurrenceattachment.DefaultGetCalendarEventExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.GetCalendarEventExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventExceptionOccurrenceAttachmentsCount(ctx, id, calendareventexceptionoccurrenceattachment.DefaultGetCalendarEventExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceAttachmentClient.ListCalendarEventExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceattachment.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventExceptionOccurrenceAttachments(ctx, id, calendareventexceptionoccurrenceattachment.DefaultListCalendarEventExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrenceAttachmentsComplete(ctx, id, calendareventexceptionoccurrenceattachment.DefaultListCalendarEventExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
