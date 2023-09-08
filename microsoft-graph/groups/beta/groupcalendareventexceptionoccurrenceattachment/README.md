
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrenceattachment` Documentation

The `groupcalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrenceattachment.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.DeleteGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceAttachmentClient.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
