
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/virtualmachinescalesetvmruncommands` Documentation

The `virtualmachinescalesetvmruncommands` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/virtualmachinescalesetvmruncommands"
```


### Client Initialization

```go
client := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVirtualMachineRunCommandID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "runCommandName")

payload := virtualmachinescalesetvmruncommands.VirtualMachineRunCommand{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineScaleSetVMRunCommandsClient.Delete`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVirtualMachineRunCommandID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "runCommandName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachineScaleSetVMRunCommandsClient.Get`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVirtualMachineRunCommandID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "runCommandName")

read, err := client.Get(ctx, id, virtualmachinescalesetvmruncommands.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachineScaleSetVMRunCommandsClient.List`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId")

// alternatively `client.List(ctx, id, virtualmachinescalesetvmruncommands.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, virtualmachinescalesetvmruncommands.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachineScaleSetVMRunCommandsClient.Update`

```go
ctx := context.TODO()
id := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVirtualMachineRunCommandID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "instanceId", "runCommandName")

payload := virtualmachinescalesetvmruncommands.VirtualMachineRunCommandUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
