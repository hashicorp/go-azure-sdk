
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `mecalendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
