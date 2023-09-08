
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventattachment` Documentation

The `usercalendargroupcalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendareventattachment.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
