
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceattachment` Documentation

The `mecalendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrenceattachment.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrenceattachment.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
