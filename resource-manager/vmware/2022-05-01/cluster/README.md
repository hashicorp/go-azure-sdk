
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/cluster` Documentation

The `cluster` SDK allows for interaction with the Azure Resource Manager Service `vmware` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/cluster"
```


### Client Initialization

```go
client := cluster.NewClusterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterClient.ListZones`

```go
ctx := context.TODO()
id := cluster.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue")

read, err := client.ListZones(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
