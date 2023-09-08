
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceattachment` Documentation

The `mecalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrenceattachment.CreateMeCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.DeleteMeCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceAttachmentClient.ListMeCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
