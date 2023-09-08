
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrenceattachment` Documentation

The `mecalendargroupcalendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
