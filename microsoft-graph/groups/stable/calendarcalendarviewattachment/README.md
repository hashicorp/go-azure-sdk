
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment` Documentation

The `calendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment"
```


### Client Initialization

```go
client := calendarcalendarviewattachment.NewCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewAttachmentClient.CreateCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewAttachment(ctx, id, payload)
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
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarviewattachment.CreateCalendarViewAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewAttachmentsUploadSession(ctx, id, payload)
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
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

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
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

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
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

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
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListCalendarViewAttachments(ctx, id, calendarcalendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewAttachmentsComplete(ctx, id, calendarcalendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
