
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/zone` Documentation

The `zone` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/zone"
```


### Client Initialization

```go
client := zone.NewZoneClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ZoneClient.ClustersListZones`

```go
ctx := context.TODO()
id := zone.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue")

read, err := client.ClustersListZones(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
