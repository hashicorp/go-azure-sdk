
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassets` Documentation

The `discoveredassets` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassets"
```


### Client Initialization

```go
client := discoveredassets.NewDiscoveredAssetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiscoveredAssetsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := discoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetName")

payload := discoveredassets.DiscoveredAsset{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveredAssetsClient.Delete`

```go
ctx := context.TODO()
id := discoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveredAssetsClient.Get`

```go
ctx := context.TODO()
id := discoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiscoveredAssetsClient.ListByResourceGroup`

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


### Example Usage: `DiscoveredAssetsClient.ListBySubscription`

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


### Example Usage: `DiscoveredAssetsClient.Update`

```go
ctx := context.TODO()
id := discoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetName")

payload := discoveredassets.DiscoveredAssetUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
