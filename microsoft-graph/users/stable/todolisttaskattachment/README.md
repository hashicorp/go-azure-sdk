
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskattachment` Documentation

The `todolisttaskattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskattachment"
```


### Client Initialization

```go
client := todolisttaskattachment.NewTodoListTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TodoListTaskAttachmentClient.CreateTodoListTaskAttachment`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

payload := todolisttaskattachment.AttachmentBase{
	// ...
}


read, err := client.CreateTodoListTaskAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.CreateTodoListTaskAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

payload := todolisttaskattachment.CreateTodoListTaskAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateTodoListTaskAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.CreateUpdateTodoListTaskAttachmentValue`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskIdAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")
var payload []byte

read, err := client.CreateUpdateTodoListTaskAttachmentValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.DeleteTodoListTaskAttachment`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskIdAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.DeleteTodoListTaskAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.GetTodoListTaskAttachment`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskIdAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetTodoListTaskAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.GetTodoListTaskAttachmentCount`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

read, err := client.GetTodoListTaskAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.GetTodoListTaskAttachmentValue`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskIdAttachmentID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetTodoListTaskAttachmentValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.ListTodoListTaskAttachments`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewUserIdTodoListIdTaskID("userIdValue", "todoTaskListIdValue", "todoTaskIdValue")

// alternatively `client.ListTodoListTaskAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListTodoListTaskAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
