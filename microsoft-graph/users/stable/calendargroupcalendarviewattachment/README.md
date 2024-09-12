
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewattachment` Documentation

The `calendargroupcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewattachment"
```


### Client Initialization

```go
client := calendargroupcalendarviewattachment.NewCalendarGroupCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.CreateCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.CreateCalendarGroupCalendarViewAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarviewattachment.CreateCalendarGroupCalendarViewAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.DeleteCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewAttachment(ctx, id, calendargroupcalendarviewattachment.DefaultDeleteCalendarGroupCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.GetCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewAttachment(ctx, id, calendargroupcalendarviewattachment.DefaultGetCalendarGroupCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.GetCalendarGroupCalendarViewAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewAttachmentsCount(ctx, id, calendargroupcalendarviewattachment.DefaultGetCalendarGroupCalendarViewAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewAttachmentClient.ListCalendarGroupCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewAttachments(ctx, id, calendargroupcalendarviewattachment.DefaultListCalendarGroupCalendarViewAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewAttachmentsComplete(ctx, id, calendargroupcalendarviewattachment.DefaultListCalendarGroupCalendarViewAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
