
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2024-09-01/volumesonpremmigration` Documentation

The `volumesonpremmigration` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2024-09-01/volumesonpremmigration"
```


### Client Initialization

```go
client := volumesonpremmigration.NewVolumesOnPremMigrationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VolumesOnPremMigrationClient.VolumesAuthorizeExternalReplication`

```go
ctx := context.TODO()
id := volumesonpremmigration.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

if err := client.VolumesAuthorizeExternalReplicationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesOnPremMigrationClient.VolumesPeerExternalCluster`

```go
ctx := context.TODO()
id := volumesonpremmigration.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

payload := volumesonpremmigration.PeerClusterForVolumeMigrationRequest{
	// ...
}


if err := client.VolumesPeerExternalClusterThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesOnPremMigrationClient.VolumesPerformReplicationTransfer`

```go
ctx := context.TODO()
id := volumesonpremmigration.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

if err := client.VolumesPerformReplicationTransferThenPoll(ctx, id); err != nil {
	// handle the error
}
```
