
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `groupcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarviewinstanceexceptionoccurrenceattachment.CreateGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
