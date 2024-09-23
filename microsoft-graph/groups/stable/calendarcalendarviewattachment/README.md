
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment` Documentation

The `calendarcalendarviewattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment"
```


### Client Initialization

```go
client := calendarcalendarviewattachment.NewCalendarCalendarViewAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewAttachmentClient.CreateCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewAttachment(ctx, id, payload, calendarcalendarviewattachment.DefaultCreateCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.CreateCalendarViewAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarviewattachment.CreateCalendarViewAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewAttachmentsUploadSession(ctx, id, payload, calendarcalendarviewattachment.DefaultCreateCalendarViewAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.DeleteCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupId", "eventId", "attachmentId")

read, err := client.DeleteCalendarViewAttachment(ctx, id, calendarcalendarviewattachment.DefaultDeleteCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.GetCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupId", "eventId", "attachmentId")

read, err := client.GetCalendarViewAttachment(ctx, id, calendarcalendarviewattachment.DefaultGetCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.GetCalendarViewAttachmentsCount`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

read, err := client.GetCalendarViewAttachmentsCount(ctx, id, calendarcalendarviewattachment.DefaultGetCalendarViewAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.ListCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

// alternatively `client.ListCalendarViewAttachments(ctx, id, calendarcalendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewAttachmentsComplete(ctx, id, calendarcalendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
