
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `mecalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrenceattachment.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
