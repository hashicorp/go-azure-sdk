
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceattachment` Documentation

The `usercalendareventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrenceattachment.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrenceattachment.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.DeleteUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceAttachmentClient.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
