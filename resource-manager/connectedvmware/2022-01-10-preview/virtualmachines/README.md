
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/virtualmachines` Documentation

The `virtualmachines` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/virtualmachines"
```


### Client Initialization

```go
client := virtualmachines.NewVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachinesClient.AssessPatches`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.AssessPatchesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Create`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachine{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Delete`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.DeleteThenPoll(ctx, id, virtualmachines.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Get`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachinesClient.InstallPatches`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachineInstallPatchesParameters{
	// ...
}


if err := client.InstallPatchesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachinesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachinesClient.Restart`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.RestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Start`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.StartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Stop`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.StopVirtualMachineOptions{
	// ...
}


if err := client.StopThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Update`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachineUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
