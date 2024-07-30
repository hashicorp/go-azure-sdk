
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


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.CreateCalendarGroupCalendarEventAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendareventattachment.CreateCalendarGroupCalendarEventAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventAttachmentCreateUploadSession(ctx, id, payload)
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

read, err := client.DeleteCalendarGroupCalendarEventAttachment(ctx, id)
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

read, err := client.GetCalendarGroupCalendarEventAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventAttachmentClient.GetCalendarGroupCalendarEventAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventattachment.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventAttachmentCount(ctx, id)
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

// alternatively `client.ListCalendarGroupCalendarEventAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
