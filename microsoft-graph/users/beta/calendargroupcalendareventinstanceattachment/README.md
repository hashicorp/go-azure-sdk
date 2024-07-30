
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceattachment` Documentation

The `calendargroupcalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventinstanceattachment.NewCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.CreateCalendarGroupCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.CreateCalendarGroupCalendarEventInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstanceattachment.CreateCalendarGroupCalendarEventInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.DeleteCalendarGroupCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.GetCalendarGroupCalendarEventInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.GetCalendarGroupCalendarEventInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceAttachmentClient.ListCalendarGroupCalendarEventInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceattachment.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
