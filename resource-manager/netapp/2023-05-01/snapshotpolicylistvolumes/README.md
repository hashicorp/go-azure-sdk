
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-05-01/snapshotpolicylistvolumes` Documentation

The `snapshotpolicylistvolumes` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-05-01/snapshotpolicylistvolumes"
```


### Client Initialization

```go
client := snapshotpolicylistvolumes.NewSnapshotPolicyListVolumesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SnapshotPolicyListVolumesClient.SnapshotPoliciesListVolumes`

```go
ctx := context.TODO()
id := snapshotpolicylistvolumes.NewSnapshotPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "snapshotPolicyName")

read, err := client.SnapshotPoliciesListVolumes(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
