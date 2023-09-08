
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrenceattachment` Documentation

The `usercalendarcalendarviewexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrenceattachment.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrenceattachment.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceAttachmentClient.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
