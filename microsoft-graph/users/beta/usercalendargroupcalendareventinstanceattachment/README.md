
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstanceattachment` Documentation

The `usercalendargroupcalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstanceattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstanceattachment.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
