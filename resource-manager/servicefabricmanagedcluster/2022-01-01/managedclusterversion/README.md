
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedclusterversion` Documentation

The `managedclusterversion` SDK allows for interaction with Azure Resource Manager `servicefabricmanagedcluster` (API Version `2022-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedclusterversion"
```


### Client Initialization

```go
client := managedclusterversion.NewManagedClusterVersionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedClusterVersionClient.Get`

```go
ctx := context.TODO()
id := managedclusterversion.NewManagedClusterVersionID("12345678-1234-9876-4563-123456789012", "location", "clusterVersion")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedClusterVersionClient.GetByEnvironment`

```go
ctx := context.TODO()
id := managedclusterversion.NewEnvironmentManagedClusterVersionID("12345678-1234-9876-4563-123456789012", "location", "clusterVersion")

read, err := client.GetByEnvironment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedClusterVersionClient.List`

```go
ctx := context.TODO()
id := managedclusterversion.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedClusterVersionClient.ListByEnvironment`

```go
ctx := context.TODO()
id := managedclusterversion.NewLocationID("12345678-1234-9876-4563-123456789012", "location")

read, err := client.ListByEnvironment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
