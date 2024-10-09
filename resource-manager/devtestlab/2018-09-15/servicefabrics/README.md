
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabrics` Documentation

The `servicefabrics` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabrics"
```


### Client Initialization

```go
client := servicefabrics.NewServiceFabricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServiceFabricsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

payload := servicefabrics.ServiceFabric{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceFabricsClient.Delete`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceFabricsClient.Get`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

read, err := client.Get(ctx, id, servicefabrics.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceFabricsClient.List`

```go
ctx := context.TODO()
id := servicefabrics.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

// alternatively `client.List(ctx, id, servicefabrics.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, servicefabrics.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServiceFabricsClient.ListApplicableSchedules`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

read, err := client.ListApplicableSchedules(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServiceFabricsClient.Start`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

if err := client.StartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceFabricsClient.Stop`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

if err := client.StopThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServiceFabricsClient.Update`

```go
ctx := context.TODO()
id := servicefabrics.NewServiceFabricID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName", "serviceFabricName")

payload := servicefabrics.UpdateResource{
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
