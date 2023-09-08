
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermessageattachment` Documentation

The `usermessageattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermessageattachment"
```


### Client Initialization

```go
client := usermessageattachment.NewUserMessageAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMessageAttachmentClient.CreateUserByIdMessageByIdAttachment`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessageattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdMessageByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageAttachmentClient.CreateUserByIdMessageByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessageattachment.CreateUserByIdMessageByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageAttachmentClient.DeleteUserByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageAttachmentClient.GetUserByIdMessageByIdAttachmentById`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue")

read, err := client.GetUserByIdMessageByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageAttachmentClient.GetUserByIdMessageByIdAttachmentCount`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.GetUserByIdMessageByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageAttachmentClient.ListUserByIdMessageByIdAttachments`

```go
ctx := context.TODO()
id := usermessageattachment.NewUserMessageID("userIdValue", "messageIdValue")

// alternatively `client.ListUserByIdMessageByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMessageByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
