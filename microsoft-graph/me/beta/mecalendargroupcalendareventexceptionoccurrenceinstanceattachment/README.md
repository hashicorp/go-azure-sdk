
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstanceattachment` Documentation

The `mecalendargroupcalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
