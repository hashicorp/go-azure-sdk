
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/todolisttaskattachment` Documentation

The `todolisttaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/todolisttaskattachment"
```


### Client Initialization

```go
client := todolisttaskattachment.NewTodoListTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TodoListTaskAttachmentClient.CreateTodoListTaskAttachment`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListIdValue", "todoTaskIdValue")

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


### Example Usage: `TodoListTaskAttachmentClient.CreateTodoListTaskAttachmentsUploadSession`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListIdValue", "todoTaskIdValue")

payload := todolisttaskattachment.CreateTodoListTaskAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateTodoListTaskAttachmentsUploadSession(ctx, id, payload)
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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.DeleteTodoListTaskAttachment(ctx, id, todolisttaskattachment.DefaultDeleteTodoListTaskAttachmentOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetTodoListTaskAttachment(ctx, id, todolisttaskattachment.DefaultGetTodoListTaskAttachmentOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetTodoListTaskAttachmentValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.GetTodoListTaskAttachmentsCount`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListIdValue", "todoTaskIdValue")

read, err := client.GetTodoListTaskAttachmentsCount(ctx, id, todolisttaskattachment.DefaultGetTodoListTaskAttachmentsCountOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListIdValue", "todoTaskIdValue")

// alternatively `client.ListTodoListTaskAttachments(ctx, id, todolisttaskattachment.DefaultListTodoListTaskAttachmentsOperationOptions())` can be used to do batched pagination
items, err := client.ListTodoListTaskAttachmentsComplete(ctx, id, todolisttaskattachment.DefaultListTodoListTaskAttachmentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TodoListTaskAttachmentClient.RemoveTodoListTaskAttachmentValue`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.RemoveTodoListTaskAttachmentValue(ctx, id, todolisttaskattachment.DefaultRemoveTodoListTaskAttachmentValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TodoListTaskAttachmentClient.SetTodoListTaskAttachmentValue`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")
var payload []byte

read, err := client.SetTodoListTaskAttachmentValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
