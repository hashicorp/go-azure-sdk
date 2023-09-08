
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventattachment` Documentation

The `mecalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventattachment"
```


### Client Initialization

```go
client := mecalendareventattachment.NewMeCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventAttachmentClient.CreateMeCalendarByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

payload := mecalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.CreateMeCalendarByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

payload := mecalendareventattachment.CreateMeCalendarByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.CreateMeCalendarEventByIdAttachment`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

payload := mecalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateMeCalendarEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.CreateMeCalendarEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

payload := mecalendareventattachment.CreateMeCalendarEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.DeleteMeCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.DeleteMeCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteMeCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.GetMeCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.GetMeCalendarByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.GetMeCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetMeCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.GetMeCalendarEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventAttachmentClient.ListMeCalendarByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

// alternatively `client.ListMeCalendarByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventAttachmentClient.ListMeCalendarEventByIdAttachments`

```go
ctx := context.TODO()
id := mecalendareventattachment.NewMeCalendarEventID("eventIdValue")

// alternatively `client.ListMeCalendarEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
