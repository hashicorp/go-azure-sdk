
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/racks` Documentation

The `racks` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/racks"
```


### Client Initialization

```go
client := racks.NewRacksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RacksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := racks.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

payload := racks.Rack{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, racks.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `RacksClient.Delete`

```go
ctx := context.TODO()
id := racks.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

if err := client.DeleteThenPoll(ctx, id, racks.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `RacksClient.Get`

```go
ctx := context.TODO()
id := racks.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RacksClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, racks.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, racks.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RacksClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, racks.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, racks.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RacksClient.Update`

```go
ctx := context.TODO()
id := racks.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

payload := racks.RackPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, racks.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
