
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceattachment` Documentation

The `usercalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrenceattachment.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceAttachmentClient.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
