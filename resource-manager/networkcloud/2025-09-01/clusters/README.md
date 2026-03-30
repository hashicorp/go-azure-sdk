
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clusters` Documentation

The `clusters` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clusters"
```


### Client Initialization

```go
client := clusters.NewClustersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClustersClient.ContinueUpdateVersion`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ClusterContinueUpdateVersionParameters{
	// ...
}


if err := client.ContinueUpdateVersionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.Cluster{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, clusters.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.Delete`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

if err := client.DeleteThenPoll(ctx, id, clusters.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.Deploy`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ClusterDeployParameters{
	// ...
}


if err := client.DeployThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.Get`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClustersClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, clusters.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, clusters.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClustersClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, clusters.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, clusters.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClustersClient.ScanRuntime`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ClusterScanRuntimeParameters{
	// ...
}


if err := client.ScanRuntimeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.Update`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ClusterPatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, clusters.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClustersClient.UpdateVersion`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ClusterUpdateVersionParameters{
	// ...
}


if err := client.UpdateVersionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
