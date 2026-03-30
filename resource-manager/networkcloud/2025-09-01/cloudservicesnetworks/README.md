
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/cloudservicesnetworks` Documentation

The `cloudservicesnetworks` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/cloudservicesnetworks"
```


### Client Initialization

```go
client := cloudservicesnetworks.NewCloudServicesNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudServicesNetworksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := cloudservicesnetworks.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

payload := cloudservicesnetworks.CloudServicesNetwork{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, cloudservicesnetworks.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `CloudServicesNetworksClient.Delete`

```go
ctx := context.TODO()
id := cloudservicesnetworks.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

if err := client.DeleteThenPoll(ctx, id, cloudservicesnetworks.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `CloudServicesNetworksClient.Get`

```go
ctx := context.TODO()
id := cloudservicesnetworks.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CloudServicesNetworksClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, cloudservicesnetworks.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, cloudservicesnetworks.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudServicesNetworksClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, cloudservicesnetworks.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, cloudservicesnetworks.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CloudServicesNetworksClient.Update`

```go
ctx := context.TODO()
id := cloudservicesnetworks.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

payload := cloudservicesnetworks.CloudServicesNetworkPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, cloudservicesnetworks.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
