
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memessageattachment` Documentation

The `memessageattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memessageattachment"
```


### Client Initialization

```go
client := memessageattachment.NewMeMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMessageAttachmentClient.CreateMeMessageByIdAttachment`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageID("messageIdValue")

payload := memessageattachment.Attachment{
	// ...
}


read, err := client.CreateMeMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageAttachmentClient.CreateMeMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageID("messageIdValue")

payload := memessageattachment.CreateMeMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageAttachmentClient.DeleteMeMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageAttachmentID("messageIdValue", "attachmentIdValue")

read, err := client.DeleteMeMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageAttachmentClient.GetMeMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageAttachmentID("messageIdValue", "attachmentIdValue")

read, err := client.GetMeMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageAttachmentClient.GetMeMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageID("messageIdValue")

read, err := client.GetMeMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageAttachmentClient.ListMeMessageByIdAttachments`

```go
ctx := context.TODO()
id := memessageattachment.NewMeMessageID("messageIdValue")

// alternatively `client.ListMeMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
