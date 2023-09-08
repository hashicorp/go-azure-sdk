
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
