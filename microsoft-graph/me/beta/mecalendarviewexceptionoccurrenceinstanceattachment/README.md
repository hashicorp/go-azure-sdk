
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `mecalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstanceattachment.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
