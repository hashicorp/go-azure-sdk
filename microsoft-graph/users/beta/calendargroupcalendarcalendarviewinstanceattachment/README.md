
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceattachment` Documentation

The `calendargroupcalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstanceattachment.NewCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstanceattachment.Attachment{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarViewInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstanceattachment.CreateCalendarGroupCalendarViewInstanceAttachmentsUploadSessionRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.DeleteCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarViewInstanceAttachment(ctx, id, calendargroupcalendarcalendarviewinstanceattachment.DefaultDeleteCalendarGroupCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarViewInstanceAttachment(ctx, id, calendargroupcalendarcalendarviewinstanceattachment.DefaultGetCalendarGroupCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarViewInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstanceAttachmentsCount(ctx, id, calendargroupcalendarcalendarviewinstanceattachment.DefaultGetCalendarGroupCalendarViewInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.ListCalendarGroupCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewInstanceAttachments(ctx, id, calendargroupcalendarcalendarviewinstanceattachment.DefaultListCalendarGroupCalendarViewInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstanceAttachmentsComplete(ctx, id, calendargroupcalendarcalendarviewinstanceattachment.DefaultListCalendarGroupCalendarViewInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
