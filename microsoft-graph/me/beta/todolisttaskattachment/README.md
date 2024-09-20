
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/todolisttaskattachment` Documentation

The `todolisttaskattachment` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/todolisttaskattachment"
```


### Client Initialization

```go
client := todolisttaskattachment.NewTodoListTaskAttachmentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TodoListTaskAttachmentClient.CreateTodoListTaskAttachment`

```go
ctx := context.TODO()
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListId", "todoTaskId")

payload := todolisttaskattachment.AttachmentBase{
	// ...
}


read, err := client.CreateTodoListTaskAttachment(ctx, id, payload, todolisttaskattachment.DefaultCreateTodoListTaskAttachmentOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListId", "todoTaskId")

payload := todolisttaskattachment.CreateTodoListTaskAttachmentsUploadSessionRequest{
	// ...
}


read, err := client.CreateTodoListTaskAttachmentsUploadSession(ctx, id, payload, todolisttaskattachment.DefaultCreateTodoListTaskAttachmentsUploadSessionOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListId", "todoTaskId", "attachmentBaseId")

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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListId", "todoTaskId", "attachmentBaseId")

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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListId", "todoTaskId", "attachmentBaseId")

read, err := client.GetTodoListTaskAttachmentValue(ctx, id, todolisttaskattachment.DefaultGetTodoListTaskAttachmentValueOperationOptions())
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
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListId", "todoTaskId")

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
id := todolisttaskattachment.NewMeTodoListIdTaskID("todoTaskListId", "todoTaskId")

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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListId", "todoTaskId", "attachmentBaseId")

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
id := todolisttaskattachment.NewMeTodoListIdTaskIdAttachmentID("todoTaskListId", "todoTaskId", "attachmentBaseId")
var payload []byte

read, err := client.SetTodoListTaskAttachmentValue(ctx, id, payload, todolisttaskattachment.DefaultSetTodoListTaskAttachmentValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
