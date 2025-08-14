
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/post` Documentation

The `post` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2025-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2025-06-30/post"
```


### Client Initialization

```go
client := post.NewPOSTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `POSTClient.FilesRead`

```go
ctx := context.TODO()
id := post.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesRead(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.FilesReadWrite`

```go
ctx := context.TODO()
id := post.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "fileName")

read, err := client.FilesReadWrite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServiceTasksCancel`

```go
ctx := context.TODO()
id := post.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "serviceTaskName")

read, err := client.ServiceTasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesCheckChildrenNameAvailability`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

payload := post.NameAvailabilityRequest{
	// ...
}


read, err := client.ServicesCheckChildrenNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := post.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := post.NameAvailabilityRequest{
	// ...
}


read, err := client.ServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesCheckStatus`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

read, err := client.ServicesCheckStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesStart`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

if err := client.ServicesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `POSTClient.ServicesStop`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

if err := client.ServicesStopThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `POSTClient.TasksCancel`

```go
ctx := context.TODO()
id := post.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "taskName")

read, err := client.TasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.TasksCommand`

```go
ctx := context.TODO()
id := post.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "taskName")

payload := post.CommandProperties{
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
