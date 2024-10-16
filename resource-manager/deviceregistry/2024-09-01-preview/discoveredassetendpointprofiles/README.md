
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassetendpointprofiles` Documentation

The `discoveredassetendpointprofiles` SDK allows for interaction with Azure Resource Manager `deviceregistry` (API Version `2024-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceregistry/2024-09-01-preview/discoveredassetendpointprofiles"
```


### Client Initialization

```go
client := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiscoveredAssetEndpointProfilesClient.CreateOrReplace`

```go
ctx := context.TODO()
id := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetEndpointProfileName")

payload := discoveredassetendpointprofiles.DiscoveredAssetEndpointProfile{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveredAssetEndpointProfilesClient.Delete`

```go
ctx := context.TODO()
id := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetEndpointProfileName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveredAssetEndpointProfilesClient.Get`

```go
ctx := context.TODO()
id := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetEndpointProfileName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiscoveredAssetEndpointProfilesClient.ListByResourceGroup`

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


### Example Usage: `DiscoveredAssetEndpointProfilesClient.ListBySubscription`

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


### Example Usage: `DiscoveredAssetEndpointProfilesClient.Update`

```go
ctx := context.TODO()
id := discoveredassetendpointprofiles.NewDiscoveredAssetEndpointProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "discoveredAssetEndpointProfileName")

payload := discoveredassetendpointprofiles.DiscoveredAssetEndpointProfileUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
