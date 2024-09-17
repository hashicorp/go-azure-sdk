
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/taskresource` Documentation

The `taskresource` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/taskresource"
```


### Client Initialization

```go
client := taskresource.NewTaskResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TaskResourceClient.TasksCancel`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TaskResourceClient.TasksCommand`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := taskresource.CommandProperties{
	// ...
}


read, err := client.TasksCommand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TaskResourceClient.TasksCreateOrUpdate`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := taskresource.ProjectTask{
	// ...
}


read, err := client.TasksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TaskResourceClient.TasksDelete`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksDelete(ctx, id, taskresource.DefaultTasksDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TaskResourceClient.TasksGet`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksGet(ctx, id, taskresource.DefaultTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TaskResourceClient.TasksUpdate`

```go
ctx := context.TODO()
id := taskresource.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := taskresource.ProjectTask{
	// ...
}


read, err := client.TasksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
