
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendareventattachment` Documentation

The `calendargroupcalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendareventattachment"
```


### Client Initialization

```go
client := calendargroupcalendareventattachment.NewCalendarGroupCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.CreateCalendarGroupCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.CreateCalendarGroupCalendarEventAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendareventattachment.CreateCalendarGroupCalendarEventAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.DeleteCalendarGroupCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarEventAttachment(ctx, id, calendargroupcalendareventattachment.DefaultDeleteCalendarGroupCalendarEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.GetCalendarGroupCalendarEventAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarEventAttachment(ctx, id, calendargroupcalendareventattachment.DefaultGetCalendarGroupCalendarEventAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.GetCalendarGroupCalendarEventAttachmentsCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventAttachmentsCount(ctx, id, calendargroupcalendareventattachment.DefaultGetCalendarGroupCalendarEventAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.ListCalendarGroupCalendarEventAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarEventAttachments(ctx, id, calendargroupcalendareventattachment.DefaultListCalendarGroupCalendarEventAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventAttachmentsComplete(ctx, id, calendargroupcalendareventattachment.DefaultListCalendarGroupCalendarEventAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
