
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrenceinstanceattachment` Documentation

The `usercalendarcalendarviewexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
