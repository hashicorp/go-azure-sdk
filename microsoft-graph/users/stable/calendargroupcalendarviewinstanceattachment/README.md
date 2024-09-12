
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewinstanceattachment` Documentation

The `calendargroupcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarviewinstanceattachment.NewCalendarGroupCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarViewInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstanceattachment.CreateCalendarGroupCalendarViewInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarViewInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.DeleteCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewInstanceAttachment(ctx, id, calendargroupcalendarviewinstanceattachment.DefaultDeleteCalendarGroupCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewInstanceAttachment(ctx, id, calendargroupcalendarviewinstanceattachment.DefaultGetCalendarGroupCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarViewInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstanceAttachmentsCount(ctx, id, calendargroupcalendarviewinstanceattachment.DefaultGetCalendarGroupCalendarViewInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceAttachmentClient.ListCalendarGroupCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceAttachments(ctx, id, calendargroupcalendarviewinstanceattachment.DefaultListCalendarGroupCalendarViewInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceAttachmentsComplete(ctx, id, calendargroupcalendarviewinstanceattachment.DefaultListCalendarGroupCalendarViewInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
