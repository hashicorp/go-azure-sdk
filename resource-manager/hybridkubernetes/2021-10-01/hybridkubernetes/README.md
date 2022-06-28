
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridkubernetes/2021-10-01/hybridkubernetes` Documentation

The `hybridkubernetes` SDK allows for interaction with the Azure Resource Manager Service `hybridkubernetes` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridkubernetes/2021-10-01/hybridkubernetes"
```


### Client Initialization

```go
client := hybridkubernetes.NewHybridKubernetesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterCreate`

```go
ctx := context.TODO()
id := hybridkubernetes.NewConnectedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := hybridkubernetes.ConnectedCluster{
	// ...
}


if err := client.ConnectedClusterCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterDelete`

```go
ctx := context.TODO()
id := hybridkubernetes.NewConnectedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

if err := client.ConnectedClusterDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterGet`

```go
ctx := context.TODO()
id := hybridkubernetes.NewConnectedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

read, err := client.ConnectedClusterGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterListByResourceGroup`

```go
ctx := context.TODO()
id := hybridkubernetes.NewResourceGroupID()

// alternatively `client.ConnectedClusterListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ConnectedClusterListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterListBySubscription`

```go
ctx := context.TODO()
id := hybridkubernetes.NewSubscriptionID()

// alternatively `client.ConnectedClusterListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ConnectedClusterListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterListClusterUserCredential`

```go
ctx := context.TODO()
id := hybridkubernetes.NewConnectedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := hybridkubernetes.ListClusterUserCredentialProperties{
	// ...
}


read, err := client.ConnectedClusterListClusterUserCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridKubernetesClient.ConnectedClusterUpdate`

```go
ctx := context.TODO()
id := hybridkubernetes.NewConnectedClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := hybridkubernetes.ConnectedClusterPatch{
	// ...
}


read, err := client.ConnectedClusterUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
