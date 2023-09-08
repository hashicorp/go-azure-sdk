
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceattachment` Documentation

The `mecalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceattachment"
```


### Client Initialization

```go
client := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.CreateMeCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.CreateMeCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstanceattachment.CreateMeCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.DeleteMeCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteMeCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.GetMeCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceAttachmentID("eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetMeCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.GetMeCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceAttachmentClient.ListMeCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
