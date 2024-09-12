
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceattachment` Documentation

The `calendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendarviewinstanceattachment.NewCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.CreateCalendarViewInstanceAttachmentsUploadSession`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstanceattachment.CreateCalendarViewInstanceAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceAttachmentsUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.DeleteCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarViewInstanceAttachment(ctx, id, calendarviewinstanceattachment.DefaultDeleteCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceIdAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarViewInstanceAttachment(ctx, id, calendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.GetCalendarViewInstanceAttachmentsCount`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstanceAttachmentsCount(ctx, id, calendarviewinstanceattachment.DefaultGetCalendarViewInstanceAttachmentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceAttachmentClient.ListCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendarviewinstanceattachment.NewUserIdCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewInstanceAttachments(ctx, id, calendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceAttachmentsComplete(ctx, id, calendarviewinstanceattachment.DefaultListCalendarViewInstanceAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
