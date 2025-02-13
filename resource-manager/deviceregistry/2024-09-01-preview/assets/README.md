
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/assets` Documentation

The `assets` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/assets"
```


### Client Initialization

```go
client := assets.NewAssetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssetsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := assets.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "assetName")

payload := assets.Asset{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AssetsClient.Delete`

```go
ctx := context.TODO()
id := assets.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "assetName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AssetsClient.Get`

```go
ctx := context.TODO()
id := assets.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "assetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssetsClient.ListByResourceGroup`

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


### Example Usage: `AssetsClient.ListBySubscription`

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


### Example Usage: `AssetsClient.Update`

```go
ctx := context.TODO()
id := assets.NewAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "assetName")

payload := assets.AssetUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
