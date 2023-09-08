
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceattachment` Documentation

The `usereventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceattachment"
```


### Client Initialization

```go
client := usereventinstanceattachment.NewUserEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventInstanceAttachmentClient.CreateUserByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceAttachmentClient.CreateUserByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventinstanceattachment.CreateUserByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceAttachmentClient.DeleteUserByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceAttachmentClient.GetUserByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceAttachmentClient.GetUserByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceAttachmentClient.ListUserByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := usereventinstanceattachment.NewUserEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
