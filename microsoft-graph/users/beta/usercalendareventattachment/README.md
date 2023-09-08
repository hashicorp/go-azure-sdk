
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventattachment` Documentation

The `usercalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventattachment"
```


### Client Initialization

```go
client := usercalendareventattachment.NewUserCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventAttachmentClient.CreateUserByIdCalendarByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.CreateUserByIdCalendarByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendareventattachment.CreateUserByIdCalendarByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.CreateUserByIdCalendarEventByIdAttachment`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.CreateUserByIdCalendarEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendareventattachment.CreateUserByIdCalendarEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.DeleteUserByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.DeleteUserByIdCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.GetUserByIdCalendarByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.GetUserByIdCalendarByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.GetUserByIdCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.GetUserByIdCalendarEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventAttachmentClient.ListUserByIdCalendarByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventAttachmentClient.ListUserByIdCalendarEventByIdAttachments`

```go
ctx := context.TODO()
id := usercalendareventattachment.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
