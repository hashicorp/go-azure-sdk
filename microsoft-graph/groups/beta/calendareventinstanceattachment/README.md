
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceattachment` Documentation

The `calendareventinstanceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceattachment"
```


### Client Initialization

```go
client := calendareventinstanceattachment.NewCalendarEventInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceAttachmentClient.CreateCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventInstanceAttachment(ctx, id, payload, calendareventinstanceattachment.DefaultCreateCalendarEventInstanceAttachmentOperationOptions())
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
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstanceattachment.CreateCalendarEventInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceAttachmentsUploadSession(ctx, id, payload, calendareventinstanceattachment.DefaultCreateCalendarEventInstanceAttachmentsUploadSessionOperationOptions())
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
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

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
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

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
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

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
id := calendareventinstanceattachment.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarEventInstanceAttachments(ctx, id, calendareventinstanceattachment.DefaultListCalendarEventInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstanceAttachmentsComplete(ctx, id, calendareventinstanceattachment.DefaultListCalendarEventInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
