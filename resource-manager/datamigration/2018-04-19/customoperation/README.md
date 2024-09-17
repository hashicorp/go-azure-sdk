
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/customoperation` Documentation

The `customoperation` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/customoperation"
```


### Client Initialization

```go
client := customoperation.NewCustomOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
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


### Example Usage: `CustomOperationClient.ServicesNestedCheckNameAvailability`

```go
ctx := context.TODO()
id := customoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

payload := customoperation.NameAvailabilityRequest{
	// ...
}


read, err := client.ServicesNestedCheckNameAvailability(ctx, id, payload)
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
