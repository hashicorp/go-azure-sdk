
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `groupcalendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
