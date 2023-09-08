
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceinstanceattachment` Documentation

The `usercalendareventexceptionoccurrenceinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceinstanceattachment"
```


### Client Initialization

```go
client := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstanceattachment.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.DeleteUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceAttachmentClient.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
