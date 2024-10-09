
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-07-01/virtualmachinescalesetvmextensions` Documentation

The `virtualmachinescalesetvmextensions` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-07-01/virtualmachinescalesetvmextensions"
```


### Client Initialization

```go
client := virtualmachinescalesetvmextensions.NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineScaleSetVMExtensionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmextensions.NewVirtualMachineExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "extensionName")

payload := virtualmachinescalesetvmextensions.VirtualMachineScaleSetVMExtension{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineScaleSetVMExtensionsClient.Delete`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmextensions.NewVirtualMachineExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "extensionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineScaleSetVMExtensionsClient.Get`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmextensions.NewVirtualMachineExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "extensionName")

read, err := client.Get(ctx, id, virtualmachinescalesetvmextensions.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineScaleSetVMExtensionsClient.List`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmextensions.NewVirtualMachineScaleSetVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId")

read, err := client.List(ctx, id, virtualmachinescalesetvmextensions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineScaleSetVMExtensionsClient.Update`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmextensions.NewVirtualMachineExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "extensionName")

payload := virtualmachinescalesetvmextensions.VirtualMachineScaleSetVMExtensionUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
