
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/mobilenetworks` Documentation

The `mobilenetworks` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-04-01/mobilenetworks"
```


### Client Initialization

```go
client := mobilenetworks.NewMobileNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MobileNetworksClient.ListByResourceGroup`

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


### Example Usage: `MobileNetworksClient.ListBySubscription`

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


### Example Usage: `MobileNetworksClient.ListSimGroups`

```go
ctx := context.TODO()
id := mobilenetworks.NewMobileNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "mobileNetworkName")

// alternatively `client.ListSimGroups(ctx, id)` can be used to do batched pagination
items, err := client.ListSimGroupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
