
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceattachment` Documentation

The `calendarviewinstanceattachment` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendarviewinstanceattachment.NewCalendarViewInstanceAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

payload := calendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachment(ctx, id, payload, calendarviewinstanceattachment.DefaultCreateCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

payload := calendarviewinstanceattachment.CreateCalendarViewInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachmentsUploadSession(ctx, id, payload, calendarviewinstanceattachment.DefaultCreateCalendarViewInstanceAttachmentsUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.DeleteCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.DeleteCalendarViewInstanceAttachment(ctx, id, calendarviewinstanceattachment.DefaultDeleteCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceIdAttachmentID("groupId", "eventId", "eventId1", "attachmentId")

read, err := client.GetCalendarViewInstanceAttachment(ctx, id, calendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarViewInstanceAttachmentsCount(ctx, id, calendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.ListCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewGroupIdCalendarViewIdInstanceID("groupId", "eventId", "eventId1")

// alternatively `client.ListCalendarViewInstanceAttachments(ctx, id, calendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceAttachmentsComplete(ctx, id, calendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
