
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceattachment` Documentation

The `calendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceattachment.NewCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstanceattachment.CreateCalendarViewInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.DeleteCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewInstanceAttachment(ctx, id, calendarcalendarviewinstanceattachment.DefaultDeleteCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarViewInstanceAttachment(ctx, id, calendarcalendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstanceAttachmentsCount(ctx, id, calendarcalendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.ListCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewInstanceAttachments(ctx, id, calendarcalendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceAttachmentsComplete(ctx, id, calendarcalendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
