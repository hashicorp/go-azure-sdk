
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewattachment` Documentation

The `mecalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewattachment"
```


### Client Initialization

```go
client := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.CreateMeCalendarByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

payload := mecalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.CreateMeCalendarByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

payload := mecalendarcalendarviewattachment.CreateMeCalendarByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.CreateMeCalendarCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

payload := mecalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.CreateMeCalendarCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

payload := mecalendarcalendarviewattachment.CreateMeCalendarCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.DeleteMeCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.DeleteMeCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.GetMeCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.GetMeCalendarByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.GetMeCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.GetMeCalendarCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.ListMeCalendarByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewAttachmentClient.ListMeCalendarCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewID("eventIdValue")

// alternatively `client.ListMeCalendarCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
