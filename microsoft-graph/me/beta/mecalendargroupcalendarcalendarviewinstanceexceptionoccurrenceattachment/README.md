
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
