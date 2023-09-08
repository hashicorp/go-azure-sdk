
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventinstanceexceptionoccurrenceattachment` Documentation

The `usercalendareventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrenceattachment.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrenceattachment.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceAttachmentClient.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
