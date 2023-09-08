
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendareventinstanceattachment` Documentation

The `usercalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendareventinstanceattachment"
```


### Client Initialization

```go
client := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstanceattachment.CreateUserByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.CreateUserByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstanceattachment.CreateUserByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.DeleteUserByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.DeleteUserByIdCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.GetUserByIdCalendarByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.GetUserByIdCalendarByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.GetUserByIdCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.GetUserByIdCalendarEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.ListUserByIdCalendarByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventInstanceAttachmentClient.ListUserByIdCalendarEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventinstanceattachment.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
