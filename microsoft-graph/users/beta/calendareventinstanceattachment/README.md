
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceattachment` Documentation

The `calendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceattachment"
```


### Client Initialization

```go
client := calendareventinstanceattachment.NewCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceAttachmentClient.CreateCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceAttachmentClient.CreateCalendarEventInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstanceattachment.CreateCalendarEventInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceAttachmentClient.DeleteCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarEventInstanceAttachment(ctx, id, calendareventinstanceattachment.DefaultDeleteCalendarEventInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceAttachmentClient.GetCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarEventInstanceAttachment(ctx, id, calendareventinstanceattachment.DefaultGetCalendarEventInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceAttachmentClient.GetCalendarEventInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventInstanceAttachmentsCount(ctx, id, calendareventinstanceattachment.DefaultGetCalendarEventInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceAttachmentClient.ListCalendarEventInstanceAttachments`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventInstanceAttachments(ctx, id, calendareventinstanceattachment.DefaultListCalendarEventInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceAttachmentsComplete(ctx, id, calendareventinstanceattachment.DefaultListCalendarEventInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
