
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewattachment` Documentation

The `mecalendargroupcalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewattachment"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarviewattachment.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewAttachmentClient.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
