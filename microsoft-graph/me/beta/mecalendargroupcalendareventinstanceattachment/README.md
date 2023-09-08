
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceattachment` Documentation

The `mecalendargroupcalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstanceattachment.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceAttachmentClient.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
