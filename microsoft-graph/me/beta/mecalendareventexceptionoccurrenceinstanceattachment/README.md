
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceinstanceattachment` Documentation

The `mecalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstanceattachment.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstanceattachment.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID("eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
