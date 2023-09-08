
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceattachment` Documentation

The `groupeventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupeventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.CreateGroupByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventexceptionoccurrenceattachment.CreateGroupByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.DeleteGroupByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.GetGroupByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.GetGroupByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventExceptionOccurrenceAttachmentClient.ListGroupByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
