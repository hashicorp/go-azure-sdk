
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewattachment` Documentation

The `usercalendargroupcalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarviewattachment.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
