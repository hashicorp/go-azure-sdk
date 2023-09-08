
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarcalendarviewinstanceattachment` Documentation

The `mecalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstanceattachment.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.CreateMeCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstanceattachment.CreateMeCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.DeleteMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.DeleteMeCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.GetMeCalendarCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.ListMeCalendarByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewInstanceAttachmentClient.ListMeCalendarCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
