
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermanagers` Documentation

The `clustermanagers` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermanagers"
```


### Client Initialization

```go
client := clustermanagers.NewClusterManagersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterManagersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := clustermanagers.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

payload := clustermanagers.ClusterManager{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, clustermanagers.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterManagersClient.Delete`

```go
ctx := context.TODO()
id := clustermanagers.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

if err := client.DeleteThenPoll(ctx, id, clustermanagers.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterManagersClient.Get`

```go
ctx := context.TODO()
id := clustermanagers.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterManagersClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, clustermanagers.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, clustermanagers.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterManagersClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, clustermanagers.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, clustermanagers.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterManagersClient.Update`

```go
ctx := context.TODO()
id := clustermanagers.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

payload := clustermanagers.ClusterManagerPatchParameters{
	// ...
}


read, err := client.Update(ctx, id, payload, clustermanagers.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
