
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewattachment` Documentation

The `calendargroupcalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewattachment.NewCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.CreateCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarviewattachment.Attachment{
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


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.CreateCalendarGroupCalendarViewAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarviewattachment.CreateCalendarGroupCalendarViewAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.DeleteCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewAttachment(ctx, id, calendargroupcalendarcalendarviewattachment.DefaultDeleteCalendarGroupCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.GetCalendarGroupCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewAttachment(ctx, id, calendargroupcalendarcalendarviewattachment.DefaultGetCalendarGroupCalendarViewAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.GetCalendarGroupCalendarViewAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewAttachmentsCount(ctx, id, calendargroupcalendarcalendarviewattachment.DefaultGetCalendarGroupCalendarViewAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.ListCalendarGroupCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewAttachments(ctx, id, calendargroupcalendarcalendarviewattachment.DefaultListCalendarGroupCalendarViewAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewAttachmentsComplete(ctx, id, calendargroupcalendarcalendarviewattachment.DefaultListCalendarGroupCalendarViewAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
