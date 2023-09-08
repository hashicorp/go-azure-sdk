
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarviewattachment` Documentation

The `mecalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarviewattachment"
```


### Client Initialization

```go
client := mecalendarviewattachment.NewMeCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewAttachmentClient.CreateMeCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewID("eventIdValue")

payload := mecalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewAttachmentClient.CreateMeCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewID("eventIdValue")

payload := mecalendarviewattachment.CreateMeCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewAttachmentClient.DeleteMeCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewAttachmentClient.GetMeCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewAttachmentClient.GetMeCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewAttachmentClient.ListMeCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarviewattachment.NewMeCalendarViewID("eventIdValue")

// alternatively `client.ListMeCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
