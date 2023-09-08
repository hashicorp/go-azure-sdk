
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarviewinstanceattachment` Documentation

The `usercalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarviewinstanceattachment"
```


### Client Initialization

```go
client := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.CreateUserByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstanceattachment.CreateUserByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.DeleteUserByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.GetUserByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.GetUserByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceAttachmentClient.ListUserByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
