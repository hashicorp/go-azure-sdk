
## `github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets` Documentation

The `fleets` SDK allows for interaction with Azure Resource Manager `azurefleet` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/azurefleet/2024-11-01/fleets"
```


### Client Initialization

```go
client := fleets.NewFleetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FleetsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := fleets.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

payload := fleets.Fleet{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FleetsClient.Delete`

```go
ctx := context.TODO()
id := fleets.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FleetsClient.Get`

```go
ctx := context.TODO()
id := fleets.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FleetsClient.ListByResourceGroup`

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


### Example Usage: `FleetsClient.ListBySubscription`

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


### Example Usage: `FleetsClient.ListVirtualMachineScaleSets`

```go
ctx := context.TODO()
id := fleets.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

// alternatively `client.ListVirtualMachineScaleSets(ctx, id)` can be used to do batched pagination
items, err := client.ListVirtualMachineScaleSetsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FleetsClient.Update`

```go
ctx := context.TODO()
id := fleets.NewFleetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "fleetName")

payload := fleets.FleetUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
