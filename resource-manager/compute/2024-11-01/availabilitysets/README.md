
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/availabilitysets` Documentation

The `availabilitysets` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/availabilitysets"
```


### Client Initialization

```go
client := availabilitysets.NewAvailabilitySetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailabilitySetsClient.CancelMigrationToVirtualMachineScaleSet`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

read, err := client.CancelMigrationToVirtualMachineScaleSet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AvailabilitySetsClient.ConvertToVirtualMachineScaleSet`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

payload := availabilitysets.ConvertToVirtualMachineScaleSetInput{
	// ...
}


if err := client.ConvertToVirtualMachineScaleSetThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AvailabilitySetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

payload := availabilitysets.AvailabilitySet{
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


### Example Usage: `AvailabilitySetsClient.Delete`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AvailabilitySetsClient.Get`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AvailabilitySetsClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AvailabilitySetsClient.ListAvailableSizes`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

// alternatively `client.ListAvailableSizes(ctx, id)` can be used to do batched pagination
items, err := client.ListAvailableSizesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AvailabilitySetsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, availabilitysets.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, availabilitysets.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AvailabilitySetsClient.StartMigrationToVirtualMachineScaleSet`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

payload := availabilitysets.MigrateToVirtualMachineScaleSetInput{
	// ...
}


read, err := client.StartMigrationToVirtualMachineScaleSet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AvailabilitySetsClient.Update`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

payload := availabilitysets.AvailabilitySetUpdate{
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


### Example Usage: `AvailabilitySetsClient.ValidateMigrationToVirtualMachineScaleSet`

```go
ctx := context.TODO()
id := commonids.NewAvailabilitySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "availabilitySetName")

payload := availabilitysets.MigrateToVirtualMachineScaleSetInput{
	// ...
}


read, err := client.ValidateMigrationToVirtualMachineScaleSet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
