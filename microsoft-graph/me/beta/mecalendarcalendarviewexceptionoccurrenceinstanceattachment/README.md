
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `mecalendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
