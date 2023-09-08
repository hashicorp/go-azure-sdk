
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `usercalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
