
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervcluster` Documentation

The `hypervcluster` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervcluster"
```


### Client Initialization

```go
client := hypervcluster.NewHyperVClusterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVClusterClient.GetAllClustersInSite`

```go
ctx := context.TODO()
id := hypervcluster.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

// alternatively `client.GetAllClustersInSite(ctx, id, hypervcluster.DefaultGetAllClustersInSiteOperationOptions())` can be used to do batched pagination
items, err := client.GetAllClustersInSiteComplete(ctx, id, hypervcluster.DefaultGetAllClustersInSiteOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVClusterClient.GetCluster`

```go
ctx := context.TODO()
id := hypervcluster.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName", "clusterName")

read, err := client.GetCluster(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HyperVClusterClient.PutCluster`

```go
ctx := context.TODO()
id := hypervcluster.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName", "clusterName")

payload := hypervcluster.HyperVCluster{
	// ...
}


read, err := client.PutCluster(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
