
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarcalendarviewinstanceattachment` Documentation

The `usercalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstanceattachment.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstanceattachment.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.DeleteUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.DeleteUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewInstanceAttachmentClient.ListUserByIdCalendarCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
