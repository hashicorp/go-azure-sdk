
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabricschedules` Documentation

The `servicefabricschedules` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabricschedules"
```


### Client Initialization

```go
client := servicefabricschedules.NewServiceFabricSchedulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServiceFabricSchedulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue", "scheduleValue")

payload := servicefabricschedules.Schedule{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceFabricSchedulesClient.Delete`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue", "scheduleValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceFabricSchedulesClient.Execute`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue", "scheduleValue")

if err := client.ExecuteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceFabricSchedulesClient.Get`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue", "scheduleValue")

read, err := client.Get(ctx, id, servicefabricschedules.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceFabricSchedulesClient.List`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue")

// alternatively `client.List(ctx, id, servicefabricschedules.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, servicefabricschedules.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServiceFabricSchedulesClient.Update`

```go
ctx := context.TODO()
id := servicefabricschedules.NewServiceFabricScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue", "serviceFabricValue", "scheduleValue")

payload := servicefabricschedules.UpdateResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
