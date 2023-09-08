
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewinstanceattachment` Documentation

The `mecalendargroupcalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstanceattachment.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
