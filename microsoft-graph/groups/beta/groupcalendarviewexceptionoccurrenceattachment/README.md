
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrenceattachment` Documentation

The `groupcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrenceattachment.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.DeleteGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceAttachmentClient.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
