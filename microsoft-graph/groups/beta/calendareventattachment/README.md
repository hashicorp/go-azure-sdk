
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventattachment` Documentation

The `calendareventattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventattachment"
```


### Client Initialization

```go
client := calendareventattachment.NewCalendarEventAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventAttachmentClient.CreateCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendareventattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventAttachment(ctx, id, payload, calendareventattachment.DefaultCreateCalendarEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.CreateCalendarEventAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendareventattachment.CreateCalendarEventAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventAttachmentsUploadSession(ctx, id, payload, calendareventattachment.DefaultCreateCalendarEventAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.DeleteCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventIdAttachmentID("groupId", "eventId", "attachmentId")

read, err := client.DeleteCalendarEventAttachment(ctx, id, calendareventattachment.DefaultDeleteCalendarEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.GetCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventIdAttachmentID("groupId", "eventId", "attachmentId")

read, err := client.GetCalendarEventAttachment(ctx, id, calendareventattachment.DefaultGetCalendarEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.GetCalendarEventAttachmentsCount`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.GetCalendarEventAttachmentsCount(ctx, id, calendareventattachment.DefaultGetCalendarEventAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.ListCalendarEventAttachments`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupId", "eventId")

// alternatively `client.ListCalendarEventAttachments(ctx, id, calendareventattachment.DefaultListCalendarEventAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventAttachmentsComplete(ctx, id, calendareventattachment.DefaultListCalendarEventAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
