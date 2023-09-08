
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usertodolisttaskattachment` Documentation

The `usertodolisttaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usertodolisttaskattachment"
```


### Client Initialization

```go
client := usertodolisttaskattachment.NewUserTodoListTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserTodoListTaskAttachmentClient.CreateUpdateUserByIdTodoListByIdTaskByIdAttachmentByIdValue`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")
var payload []byte

read, err := client.CreateUpdateUserByIdTodoListByIdTaskByIdAttachmentByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.CreateUserByIdTodoListByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

payload := usertodolisttaskattachment.AttachmentBase{
	// ...
}


read, err := client.CreateUserByIdTodoListByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.CreateUserByIdTodoListByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

payload := usertodolisttaskattachment.CreateUserByIdTodoListByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdTodoListByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.DeleteUserByIdTodoListByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.DeleteUserByIdTodoListByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.GetUserByIdTodoListByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetUserByIdTodoListByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.GetUserByIdTodoListByIdTaskByIdAttachmentByIdValue`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetUserByIdTodoListByIdTaskByIdAttachmentByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.GetUserByIdTodoListByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

read, err := client.GetUserByIdTodoListByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserTodoListTaskAttachmentClient.ListUserByIdTodoListByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := usertodolisttaskattachment.NewUserTodoListTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

// alternatively `client.ListUserByIdTodoListByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdTodoListByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
