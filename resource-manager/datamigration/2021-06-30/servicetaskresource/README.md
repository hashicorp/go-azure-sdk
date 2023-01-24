
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/servicetaskresource` Documentation

The `servicetaskresource` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/servicetaskresource"
```


### Client Initialization

```go
client := servicetaskresource.NewServiceTaskResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServiceTaskResourceClient.ServiceTasksCancel`

```go
ctx := context.TODO()
id := servicetaskresource.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceTaskResourceClient.ServiceTasksCreateOrUpdate`

```go
ctx := context.TODO()
id := servicetaskresource.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

payload := servicetaskresource.ProjectTask{
	// ...
}


read, err := client.ServiceTasksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceTaskResourceClient.ServiceTasksDelete`

```go
ctx := context.TODO()
id := servicetaskresource.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksDelete(ctx, id, servicetaskresource.DefaultServiceTasksDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceTaskResourceClient.ServiceTasksGet`

```go
ctx := context.TODO()
id := servicetaskresource.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksGet(ctx, id, servicetaskresource.DefaultServiceTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceTaskResourceClient.ServiceTasksUpdate`

```go
ctx := context.TODO()
id := servicetaskresource.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

payload := servicetaskresource.ProjectTask{
	// ...
}


read, err := client.ServiceTasksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
