
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventinstanceexceptionoccurrenceattachment` Documentation

The `groupeventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventinstanceexceptionoccurrenceattachment.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.DeleteGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceExceptionOccurrenceAttachmentClient.ListGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
