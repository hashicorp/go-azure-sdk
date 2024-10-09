
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/virtualmachineschedules` Documentation

The `virtualmachineschedules` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/virtualmachineschedules"
```


### Client Initialization

```go
client := virtualmachineschedules.NewVirtualMachineSchedulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineSchedulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName", "scheduleName")

payload := virtualmachineschedules.Schedule{
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


### Example Usage: `VirtualMachineSchedulesClient.Delete`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName", "scheduleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineSchedulesClient.Execute`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName", "scheduleName")

if err := client.ExecuteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineSchedulesClient.Get`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName", "scheduleName")

read, err := client.Get(ctx, id, virtualmachineschedules.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineSchedulesClient.List`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName")

// alternatively `client.List(ctx, id, virtualmachineschedules.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, virtualmachineschedules.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachineSchedulesClient.Update`

```go
ctx := context.TODO()
id := virtualmachineschedules.NewVirtualMachineScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "virtualMachineName", "scheduleName")

payload := virtualmachineschedules.UpdateResource{
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
