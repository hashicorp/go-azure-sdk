
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `calendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewinstanceexceptionoccurrenceattachment.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultDeleteCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarViewInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetCalendarViewInstanceExceptionOccurrenceAttachment(ctx, id, calendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarViewInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetCalendarViewInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewInstanceExceptionOccurrenceAttachmentsCount(ctx, id, calendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultGetCalendarViewInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListCalendarViewInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListCalendarViewInstanceExceptionOccurrenceAttachments(ctx, id, calendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, calendarcalendarviewinstanceexceptionoccurrenceattachment.DefaultListCalendarViewInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
