
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/virtualmachines` Documentation

The `virtualmachines` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/virtualmachines"
```


### Client Initialization

```go
client := virtualmachines.NewVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachinesClient.AssignRelay`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachineAssignRelayParameters{
	// ...
}


if err := client.AssignRelayThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachine{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, virtualmachines.DefaultCreateOrUpdateOperationOptions()); err != nil {
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


### Example Usage: `VirtualMachinesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, virtualmachines.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, virtualmachines.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachinesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, virtualmachines.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, virtualmachines.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachinesClient.PowerOff`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachinePowerOffParameters{
	// ...
}


if err := client.PowerOffThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualMachinesClient.Reimage`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.ReimageThenPoll(ctx, id); err != nil {
	// handle the error
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


### Example Usage: `VirtualMachinesClient.Update`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := virtualmachines.VirtualMachinePatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, virtualmachines.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
