
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceattachment` Documentation

The `mecalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrenceattachment.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrenceattachment.CreateMeCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.DeleteMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.DeleteMeCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceAttachmentClient.ListMeCalendarEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
