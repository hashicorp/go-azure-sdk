
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventinstanceattachment` Documentation

The `mecalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventinstanceattachment"
```


### Client Initialization

```go
client := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.CreateMeCalendarByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

payload := mecalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.CreateMeCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

payload := mecalendareventinstanceattachment.CreateMeCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.CreateMeCalendarEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

payload := mecalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.CreateMeCalendarEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

payload := mecalendareventinstanceattachment.CreateMeCalendarEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.DeleteMeCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.DeleteMeCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.GetMeCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.GetMeCalendarByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.GetMeCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentID("calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.GetMeCalendarEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.ListMeCalendarByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventInstanceAttachmentClient.ListMeCalendarEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventinstanceattachment.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
