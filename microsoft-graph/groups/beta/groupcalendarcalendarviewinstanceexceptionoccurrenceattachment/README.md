
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `groupcalendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
