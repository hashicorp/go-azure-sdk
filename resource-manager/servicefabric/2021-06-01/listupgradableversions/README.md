
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/listupgradableversions` Documentation

The `listupgradableversions` SDK allows for interaction with Azure Resource Manager `servicefabric` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabric/2021-06-01/listupgradableversions"
```


### Client Initialization

```go
client := listupgradableversions.NewListUpgradableVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListUpgradableVersionsClient.ClustersListUpgradableVersions`

```go
ctx := context.TODO()
id := listupgradableversions.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := listupgradableversions.UpgradableVersionsDescription{
	// ...
}


read, err := client.ClustersListUpgradableVersions(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
