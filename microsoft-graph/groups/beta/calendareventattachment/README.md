
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventattachment` Documentation

The `calendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventattachment"
```


### Client Initialization

```go
client := calendareventattachment.NewCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventAttachmentClient.CreateCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

payload := calendareventattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarEventAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.CreateCalendarEventAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

payload := calendareventattachment.CreateCalendarEventAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarEventAttachmentCreateUploadSession(ctx, id, payload)
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
id := calendareventattachment.NewGroupIdCalendarEventIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarEventAttachment(ctx, id)
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
id := calendareventattachment.NewGroupIdCalendarEventIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarEventAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventAttachmentClient.GetCalendarEventAttachmentCount`

```go
ctx := context.TODO()
id := calendareventattachment.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarEventAttachmentCount(ctx, id)
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
id := calendareventattachment.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListCalendarEventAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
