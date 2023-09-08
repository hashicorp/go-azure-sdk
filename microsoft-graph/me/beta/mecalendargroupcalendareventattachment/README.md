
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventattachment` Documentation

The `mecalendargroupcalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendareventattachment.CreateMeCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.GetMeCalendarGroupByIdCalendarByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventAttachmentClient.ListMeCalendarGroupByIdCalendarByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
