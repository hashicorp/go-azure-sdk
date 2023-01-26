
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/clusterversion` Documentation

The `clusterversion` SDK allows for interaction with the Azure Resource Manager Service `servicefabric` (API Version `2021-06-01`).

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
id := clusterversion.NewClusterVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "clusterVersionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterVersionClient.GetByEnvironment`

```go
ctx := context.TODO()
id := clusterversion.NewEnvironmentClusterVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "Linux", "clusterVersionValue")

read, err := client.GetByEnvironment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterVersionClient.List`

```go
ctx := context.TODO()
id := clusterversion.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterVersionClient.ListByEnvironment`

```go
ctx := context.TODO()
id := clusterversion.NewEnvironmentID("12345678-1234-9876-4563-123456789012", "locationValue", "Linux")

read, err := client.ListByEnvironment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
