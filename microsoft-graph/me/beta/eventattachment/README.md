
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventattachment` Documentation

The `eventattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/eventattachment"
```


### Client Initialization

```go
client := eventattachment.NewEventAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventAttachmentClient.CreateEventAttachment`

```go
ctx := context.TODO()
id := eventattachment.NewMeEventID("eventId")

payload := eventattachment.Attachment{
	// ...
}


read, err := client.CreateEventAttachment(ctx, id, payload, eventattachment.DefaultCreateEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.CreateEventAttachmentsUploadSession`

```go
ctx := context.TODO()
id := eventattachment.NewMeEventID("eventId")

payload := eventattachment.CreateEventAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateEventAttachmentsUploadSession(ctx, id, payload, eventattachment.DefaultCreateEventAttachmentsUploadSessionOperationOptions())
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
id := eventattachment.NewMeEventIdAttachmentID("eventId", "attachmentId")

read, err := client.DeleteEventAttachment(ctx, id, eventattachment.DefaultDeleteEventAttachmentOperationOptions())
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
id := eventattachment.NewMeEventIdAttachmentID("eventId", "attachmentId")

read, err := client.GetEventAttachment(ctx, id, eventattachment.DefaultGetEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventAttachmentClient.GetEventAttachmentsCount`

```go
ctx := context.TODO()
id := eventattachment.NewMeEventID("eventId")

read, err := client.GetEventAttachmentsCount(ctx, id, eventattachment.DefaultGetEventAttachmentsCountOperationOptions())
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
id := eventattachment.NewMeEventID("eventId")

// alternatively `client.ListEventAttachments(ctx, id, eventattachment.DefaultListEventAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventAttachmentsComplete(ctx, id, eventattachment.DefaultListEventAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
