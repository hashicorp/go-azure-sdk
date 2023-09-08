
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstanceattachment` Documentation

The `usercalendargroupcalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
