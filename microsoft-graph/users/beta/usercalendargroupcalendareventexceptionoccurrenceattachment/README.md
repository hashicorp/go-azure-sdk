
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceattachment` Documentation

The `usercalendargroupcalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrenceattachment.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
