
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewattachment` Documentation

The `usercalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewattachment"
```


### Client Initialization

```go
client := usercalendarviewattachment.NewUserCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewAttachmentClient.CreateUserByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewAttachmentClient.CreateUserByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarviewattachment.CreateUserByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewAttachmentClient.DeleteUserByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewAttachmentClient.GetUserByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewAttachmentClient.GetUserByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewAttachmentClient.ListUserByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarviewattachment.NewUserCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
