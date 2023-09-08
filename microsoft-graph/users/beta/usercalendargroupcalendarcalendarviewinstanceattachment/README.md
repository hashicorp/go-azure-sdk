
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstanceattachment` Documentation

The `usercalendargroupcalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstanceattachment.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
