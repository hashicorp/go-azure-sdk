
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceattachment` Documentation

The `mecalendargroupcalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrenceattachment.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
