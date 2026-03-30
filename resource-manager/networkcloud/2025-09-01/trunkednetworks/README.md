
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/trunkednetworks` Documentation

The `trunkednetworks` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/trunkednetworks"
```


### Client Initialization

```go
client := trunkednetworks.NewTrunkedNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TrunkedNetworksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := trunkednetworks.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

payload := trunkednetworks.TrunkedNetwork{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, trunkednetworks.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrunkedNetworksClient.Delete`

```go
ctx := context.TODO()
id := trunkednetworks.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

if err := client.DeleteThenPoll(ctx, id, trunkednetworks.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `TrunkedNetworksClient.Get`

```go
ctx := context.TODO()
id := trunkednetworks.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TrunkedNetworksClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, trunkednetworks.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, trunkednetworks.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TrunkedNetworksClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, trunkednetworks.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, trunkednetworks.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TrunkedNetworksClient.Update`

```go
ctx := context.TODO()
id := trunkednetworks.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

payload := trunkednetworks.TrunkedNetworkPatchParameters{
	// ...
}


read, err := client.Update(ctx, id, payload, trunkednetworks.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
