
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/tasks` Documentation

The `tasks` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/tasks"
```


### Client Initialization

```go
client := tasks.NewTasksClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `TasksClient.TaskAdd`

```go
ctx := context.TODO()
id := tasks.NewJobID("jobId")

payload := tasks.TaskAddParameter{
	// ...
}


read, err := client.TaskAdd(ctx, id, payload, tasks.DefaultTaskAddOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskAddCollection`

```go
ctx := context.TODO()
id := tasks.NewJobID("jobId")

payload := tasks.TaskAddCollectionParameter{
	// ...
}


read, err := client.TaskAddCollection(ctx, id, payload, tasks.DefaultTaskAddCollectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskDelete`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

read, err := client.TaskDelete(ctx, id, tasks.DefaultTaskDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskGet`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

read, err := client.TaskGet(ctx, id, tasks.DefaultTaskGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskList`

```go
ctx := context.TODO()
id := tasks.NewJobID("jobId")

// alternatively `client.TaskList(ctx, id, tasks.DefaultTaskListOperationOptions())` can be used to do batched pagination
items, err := client.TaskListComplete(ctx, id, tasks.DefaultTaskListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TasksClient.TaskListSubtasks`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

read, err := client.TaskListSubtasks(ctx, id, tasks.DefaultTaskListSubtasksOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskReactivate`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

read, err := client.TaskReactivate(ctx, id, tasks.DefaultTaskReactivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskTerminate`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

read, err := client.TaskTerminate(ctx, id, tasks.DefaultTaskTerminateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TasksClient.TaskUpdate`

```go
ctx := context.TODO()
id := tasks.NewTaskID("jobId", "taskId")

payload := tasks.TaskUpdateParameter{
	// ...
}


read, err := client.TaskUpdate(ctx, id, payload, tasks.DefaultTaskUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
