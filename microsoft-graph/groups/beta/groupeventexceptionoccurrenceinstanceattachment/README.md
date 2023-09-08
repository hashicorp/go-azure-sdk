
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceinstanceattachment` Documentation

The `groupeventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupeventexceptionoccurrenceinstanceattachment.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.DeleteGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceInstanceAttachmentClient.ListGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
