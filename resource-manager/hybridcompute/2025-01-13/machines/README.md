
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machines` Documentation

The `machines` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machines"
```


### Client Initialization

```go
client := machines.NewMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachinesClient.AssessPatches`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

if err := client.AssessPatchesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MachinesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := machines.Machine{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, machines.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.Delete`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MachinesClient.Get`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

read, err := client.Get(ctx, id, machines.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.InstallPatches`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := machines.MachineInstallPatchesParameters{
	// ...
}


if err := client.InstallPatchesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MachinesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, machines.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, machines.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachinesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachinesClient.Update`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := machines.MachineUpdate{
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
