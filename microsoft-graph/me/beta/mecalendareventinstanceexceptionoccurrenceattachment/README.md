
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventinstanceexceptionoccurrenceattachment` Documentation

The `mecalendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrenceattachment.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrenceattachment.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
