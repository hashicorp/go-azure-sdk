
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarviewattachment` Documentation

The `calendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarviewattachment"
```


### Client Initialization

```go
client := calendarviewattachment.NewCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewAttachmentClient.CreateCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewID("eventIdValue")

payload := calendarviewattachment.Attachment{
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


### Example Usage: `CalendarViewAttachmentClient.CreateCalendarViewAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewID("eventIdValue")

payload := calendarviewattachment.CreateCalendarViewAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarViewAttachmentClient.DeleteCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewIdAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarViewAttachment(ctx, id, calendarviewattachment.DefaultDeleteCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewAttachmentClient.GetCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewIdAttachmentID("eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarViewAttachment(ctx, id, calendarviewattachment.DefaultGetCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewAttachmentClient.GetCalendarViewAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewID("eventIdValue")

read, err := client.GetCalendarViewAttachmentsCount(ctx, id, calendarviewattachment.DefaultGetCalendarViewAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewAttachmentClient.ListCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendarviewattachment.NewMeCalendarViewID("eventIdValue")

// alternatively `client.ListCalendarViewAttachments(ctx, id, calendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewAttachmentsComplete(ctx, id, calendarviewattachment.DefaultListCalendarViewAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
