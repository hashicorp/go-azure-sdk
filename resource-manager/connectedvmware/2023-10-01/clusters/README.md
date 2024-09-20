
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/clusters` Documentation

The `clusters` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/clusters"
```


### Client Initialization

```go
client := clusters.NewClustersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClustersClient.Create`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.Cluster{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
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


### Example Usage: `ClustersClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClustersClient.ListByResourceGroup`

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


### Example Usage: `ClustersClient.Update`

```go
ctx := context.TODO()
id := clusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := clusters.ResourcePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
