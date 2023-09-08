
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/metodolisttaskattachment` Documentation

The `metodolisttaskattachment` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/metodolisttaskattachment"
```


### Client Initialization

```go
client := metodolisttaskattachment.NewMeTodoListTaskAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeTodoListTaskAttachmentClient.CreateMeTodoListByIdTaskByIdAttachment`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue")

payload := metodolisttaskattachment.AttachmentBase{
	// ...
}


read, err := client.CreateMeTodoListByIdTaskByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.CreateMeTodoListByIdTaskByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue")

payload := metodolisttaskattachment.CreateMeTodoListByIdTaskByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateMeTodoListByIdTaskByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.CreateUpdateMeTodoListByIdTaskByIdAttachmentByIdValue`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")
var payload []byte

read, err := client.CreateUpdateMeTodoListByIdTaskByIdAttachmentByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.DeleteMeTodoListByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.DeleteMeTodoListByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.GetMeTodoListByIdTaskByIdAttachmentById`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetMeTodoListByIdTaskByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.GetMeTodoListByIdTaskByIdAttachmentByIdValue`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskAttachmentID("todoTaskListIdValue", "todoTaskIdValue", "attachmentBaseIdValue")

read, err := client.GetMeTodoListByIdTaskByIdAttachmentByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.GetMeTodoListByIdTaskByIdAttachmentCount`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue")

read, err := client.GetMeTodoListByIdTaskByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeTodoListTaskAttachmentClient.ListMeTodoListByIdTaskByIdAttachments`

```go
ctx := context.TODO()
id := metodolisttaskattachment.NewMeTodoListTaskID("todoTaskListIdValue", "todoTaskIdValue")

// alternatively `client.ListMeTodoListByIdTaskByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListMeTodoListByIdTaskByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
