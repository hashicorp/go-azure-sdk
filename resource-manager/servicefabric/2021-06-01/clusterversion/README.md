
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/clusterversion` Documentation

The `clusterversion` SDK allows for interaction with Azure Resource Manager `servicefabric` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/clusterversion"
```


### Client Initialization

```go
client := clusterversion.NewClusterVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterVersionClient.Get`

```go
ctx := context.TODO()
id := clusterversion.NewClusterVersionID("12345678-1234-9876-4563-123456789012", "locationName", "clusterVersionName")

// alternatively `client.Get(ctx, id)` can be used to do batched pagination
items, err := client.GetComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterVersionClient.GetByEnvironment`

```go
ctx := context.TODO()
id := clusterversion.NewEnvironmentClusterVersionID("12345678-1234-9876-4563-123456789012", "locationName", "Linux", "clusterVersionName")

// alternatively `client.GetByEnvironment(ctx, id)` can be used to do batched pagination
items, err := client.GetByEnvironmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterVersionClient.List`

```go
ctx := context.TODO()
id := clusterversion.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterVersionClient.ListByEnvironment`

```go
ctx := context.TODO()
id := clusterversion.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "locationName", "Linux")

// alternatively `client.ListByEnvironment(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnvironmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
