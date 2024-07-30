
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceattachment` Documentation

The `eventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceattachment"
```


### Client Initialization

```go
client := eventinstanceattachment.NewEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceAttachmentClient.CreateEventInstanceAttachment`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateEventInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.CreateEventInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstanceattachment.CreateEventInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateEventInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.DeleteEventInstanceAttachment`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteEventInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.GetEventInstanceAttachment`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetEventInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.GetEventInstanceAttachmentCount`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.ListEventInstanceAttachments`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListEventInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListEventInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
