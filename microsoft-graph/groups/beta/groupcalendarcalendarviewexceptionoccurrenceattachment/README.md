
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceattachment` Documentation

The `groupcalendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrenceattachment.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
