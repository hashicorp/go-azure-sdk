
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarcalendarviewinstanceattachment` Documentation

The `calendargroupcalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstanceattachment.NewCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateCalendarGroupCalendarCalendarViewInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstanceattachment.CreateCalendarGroupCalendarCalendarViewInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.DeleteCalendarGroupCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarCalendarViewInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdAttachmentID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetCalendarGroupCalendarCalendarViewInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceAttachmentClient.ListCalendarGroupCalendarCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupIdCalendarIdCalendarViewIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarCalendarViewInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
