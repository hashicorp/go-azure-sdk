
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrenceattachment` Documentation

The `eventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := eventinstanceexceptionoccurrenceattachment.NewEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.CreateEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.CreateEventInstanceExceptionOccurrenceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrenceattachment.CreateEventInstanceExceptionOccurrenceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventInstanceExceptionOccurrenceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.DeleteEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteEventInstanceExceptionOccurrenceAttachment(ctx, id, eventinstanceexceptionoccurrenceattachment.DefaultDeleteEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.GetEventInstanceExceptionOccurrenceAttachment`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetEventInstanceExceptionOccurrenceAttachment(ctx, id, eventinstanceexceptionoccurrenceattachment.DefaultGetEventInstanceExceptionOccurrenceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.GetEventInstanceExceptionOccurrenceAttachmentsCount`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetEventInstanceExceptionOccurrenceAttachmentsCount(ctx, id, eventinstanceexceptionoccurrenceattachment.DefaultGetEventInstanceExceptionOccurrenceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceAttachmentClient.ListEventInstanceExceptionOccurrenceAttachments`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrenceattachment.NewGroupIdEventIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListEventInstanceExceptionOccurrenceAttachments(ctx, id, eventinstanceexceptionoccurrenceattachment.DefaultListEventInstanceExceptionOccurrenceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventInstanceExceptionOccurrenceAttachmentsComplete(ctx, id, eventinstanceexceptionoccurrenceattachment.DefaultListEventInstanceExceptionOccurrenceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
