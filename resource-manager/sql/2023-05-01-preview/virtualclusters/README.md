
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/virtualclusters` Documentation

The `virtualclusters` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/virtualclusters"
```


### Client Initialization

```go
client := virtualclusters.NewVirtualClustersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualClustersClient.Delete`

```go
ctx := context.TODO()
id := virtualclusters.NewVirtualClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualClusterValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualClustersClient.Get`

```go
ctx := context.TODO()
id := virtualclusters.NewVirtualClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualClusterValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualClustersClient.List`

```go
ctx := context.TODO()
id := virtualclusters.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualClustersClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := virtualclusters.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualClustersClient.Update`

```go
ctx := context.TODO()
id := virtualclusters.NewVirtualClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualClusterValue")

payload := virtualclusters.VirtualClusterUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualClustersClient.UpdateDnsServers`

```go
ctx := context.TODO()
id := virtualclusters.NewVirtualClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualClusterValue")

if err := client.UpdateDnsServersThenPoll(ctx, id); err != nil {
	// handle the error
}
```
