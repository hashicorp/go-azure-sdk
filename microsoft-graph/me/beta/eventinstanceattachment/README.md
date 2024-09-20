
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventinstanceattachment` Documentation

The `eventinstanceattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventinstanceattachment"
```


### Client Initialization

```go
client := eventinstanceattachment.NewEventInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceAttachmentClient.CreateEventInstanceAttachment`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewMeEventIdInstanceID("eventId", "eventId1")

payload := eventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateEventInstanceAttachment(ctx, id, payload, eventinstanceattachment.DefaultCreateEventInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.CreateEventInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewMeEventIdInstanceID("eventId", "eventId1")

payload := eventinstanceattachment.CreateEventInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventInstanceAttachmentsUploadSession(ctx, id, payload, eventinstanceattachment.DefaultCreateEventInstanceAttachmentsUploadSessionOperationOptions())
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
id := eventinstanceattachment.NewMeEventIdInstanceIdAttachmentID("eventId", "eventId1", "attachmentId")

read, err := client.DeleteEventInstanceAttachment(ctx, id, eventinstanceattachment.DefaultDeleteEventInstanceAttachmentOperationOptions())
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
id := eventinstanceattachment.NewMeEventIdInstanceIdAttachmentID("eventId", "eventId1", "attachmentId")

read, err := client.GetEventInstanceAttachment(ctx, id, eventinstanceattachment.DefaultGetEventInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceAttachmentClient.GetEventInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := eventinstanceattachment.NewMeEventIdInstanceID("eventId", "eventId1")

read, err := client.GetEventInstanceAttachmentsCount(ctx, id, eventinstanceattachment.DefaultGetEventInstanceAttachmentsCountOperationOptions())
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
id := eventinstanceattachment.NewMeEventIdInstanceID("eventId", "eventId1")

// alternatively `client.ListEventInstanceAttachments(ctx, id, eventinstanceattachment.DefaultListEventInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventInstanceAttachmentsComplete(ctx, id, eventinstanceattachment.DefaultListEventInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
