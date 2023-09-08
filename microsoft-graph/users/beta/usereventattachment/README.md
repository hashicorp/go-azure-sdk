
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventattachment` Documentation

The `usereventattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventattachment"
```


### Client Initialization

```go
client := usereventattachment.NewUserEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventAttachmentClient.CreateUserByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventID("userIdValue", "eventIdValue")

payload := usereventattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventAttachmentClient.CreateUserByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventID("userIdValue", "eventIdValue")

payload := usereventattachment.CreateUserByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventAttachmentClient.DeleteUserByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventAttachmentClient.GetUserByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventAttachmentID("userIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetUserByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventAttachmentClient.GetUserByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventAttachmentClient.ListUserByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := usereventattachment.NewUserEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
