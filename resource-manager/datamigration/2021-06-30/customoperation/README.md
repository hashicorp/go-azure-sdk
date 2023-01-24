
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/customoperation` Documentation

The `customoperation` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/customoperation"
```


### Client Initialization

```go
client := customoperation.NewCustomOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomOperationClient.ServiceTasksCancel`

```go
ctx := context.TODO()
id := customoperation.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomOperationClient.ServicesCheckChildrenNameAvailability`

```go
ctx := context.TODO()
id := customoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

payload := customoperation.NameAvailabilityRequest{
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


### Example Usage: `CustomOperationClient.ServicesCheckStatus`

```go
ctx := context.TODO()
id := customoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

read, err := client.ServicesCheckStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomOperationClient.ServicesStart`

```go
ctx := context.TODO()
id := customoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

if err := client.ServicesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CustomOperationClient.ServicesStop`

```go
ctx := context.TODO()
id := customoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

if err := client.ServicesStopThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CustomOperationClient.TasksCancel`

```go
ctx := context.TODO()
id := customoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomOperationClient.TasksCommand`

```go
ctx := context.TODO()
id := customoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := customoperation.CommandProperties{
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
