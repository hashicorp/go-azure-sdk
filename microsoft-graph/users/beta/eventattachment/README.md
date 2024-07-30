
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventattachment` Documentation

The `eventattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventattachment"
```


### Client Initialization

```go
client := eventattachment.NewEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventAttachmentClient.CreateEventAttachment`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventID("userIdValue", "eventIdValue")

payload := eventattachment.Attachment{
	// ...
}


read, err := client.CreateEventAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.CreateEventAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventID("userIdValue", "eventIdValue")

payload := eventattachment.CreateEventAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateEventAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.DeleteEventAttachment`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventIdAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteEventAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.GetEventAttachment`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventIdAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetEventAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.GetEventAttachmentCount`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventID("userIdValue", "eventIdValue")

read, err := client.GetEventAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.ListEventAttachments`

```go
ctx := context.TODO()
id := eventattachment.NewUserIdEventID("userIdValue", "eventIdValue")

// alternatively `client.ListEventAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListEventAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
