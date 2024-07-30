
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


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.CreateCalendarGroupCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.CreateCalendarGroupCalendarCalendarViewAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarviewattachment.CreateCalendarGroupCalendarCalendarViewAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.DeleteCalendarGroupCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarGroupCalendarCalendarViewAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.GetCalendarGroupCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.GetCalendarGroupCalendarCalendarViewAttachmentCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewAttachmentClient.ListCalendarGroupCalendarCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewattachment.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarCalendarViewAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
