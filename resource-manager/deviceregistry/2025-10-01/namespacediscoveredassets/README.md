
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscoveredassets` Documentation

The `namespacediscoveredassets` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2025-10-01/namespacediscoveredassets"
```


### Client Initialization

```go
client := namespacediscoveredassets.NewNamespaceDiscoveredAssetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NamespaceDiscoveredAssetsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := namespacediscoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredAssetName")

payload := namespacediscoveredassets.NamespaceDiscoveredAsset{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDiscoveredAssetsClient.Delete`

```go
ctx := context.TODO()
id := namespacediscoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredAssetName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NamespaceDiscoveredAssetsClient.Get`

```go
ctx := context.TODO()
id := namespacediscoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredAssetName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespaceDiscoveredAssetsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := namespacediscoveredassets.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NamespaceDiscoveredAssetsClient.Update`

```go
ctx := context.TODO()
id := namespacediscoveredassets.NewDiscoveredAssetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "discoveredAssetName")

payload := namespacediscoveredassets.NamespaceDiscoveredAssetUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
