
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/volumesrelocation` Documentation

The `volumesrelocation` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/volumesrelocation"
```


### Client Initialization

```go
client := volumesrelocation.NewVolumesRelocationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VolumesRelocationClient.VolumesFinalizeRelocation`

```go
ctx := context.TODO()
id := volumesrelocation.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

if err := client.VolumesFinalizeRelocationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesRelocationClient.VolumesRelocate`

```go
ctx := context.TODO()
id := volumesrelocation.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

payload := volumesrelocation.RelocateVolumeRequest{
	// ...
}


if err := client.VolumesRelocateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesRelocationClient.VolumesRevertRelocation`

```go
ctx := context.TODO()
id := volumesrelocation.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

if err := client.VolumesRevertRelocationThenPoll(ctx, id); err != nil {
	// handle the error
}
```
