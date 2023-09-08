
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarcalendarviewattachment` Documentation

The `usercalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarcalendarviewattachment"
```


### Client Initialization

```go
client := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarcalendarviewattachment.CreateUserByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.CreateUserByIdCalendarCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarcalendarviewattachment.CreateUserByIdCalendarCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.DeleteUserByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.DeleteUserByIdCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.GetUserByIdCalendarByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.GetUserByIdCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.GetUserByIdCalendarCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.ListUserByIdCalendarByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewAttachmentClient.ListUserByIdCalendarCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
