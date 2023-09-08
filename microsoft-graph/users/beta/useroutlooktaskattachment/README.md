
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskattachment` Documentation

The `useroutlooktaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useroutlooktaskattachment"
```


### Client Initialization

```go
client := useroutlooktaskattachment.NewUserOutlookTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOutlookTaskAttachmentClient.CreateUserByIdOutlookTaskByIdAttachment`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue")

payload := useroutlooktaskattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskAttachmentClient.CreateUserByIdOutlookTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue")

payload := useroutlooktaskattachment.CreateUserByIdOutlookTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdOutlookTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskAttachmentClient.DeleteUserByIdOutlookTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskAttachmentID("userIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.DeleteUserByIdOutlookTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskAttachmentClient.GetUserByIdOutlookTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskAttachmentID("userIdValue", "outlookTaskIdValue", "attachmentIdValue")

read, err := client.GetUserByIdOutlookTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskAttachmentClient.GetUserByIdOutlookTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue")

read, err := client.GetUserByIdOutlookTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOutlookTaskAttachmentClient.ListUserByIdOutlookTaskByIdAttachments`

```go
ctx := context.TODO()
id := useroutlooktaskattachment.NewUserOutlookTaskID("userIdValue", "outlookTaskIdValue")

// alternatively `client.ListUserByIdOutlookTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOutlookTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
