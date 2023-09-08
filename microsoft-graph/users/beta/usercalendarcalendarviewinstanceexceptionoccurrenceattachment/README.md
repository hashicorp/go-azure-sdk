
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstanceexceptionoccurrenceattachment` Documentation

The `usercalendarcalendarviewinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
