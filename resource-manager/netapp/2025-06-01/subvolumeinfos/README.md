
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/subvolumeinfos` Documentation

The `subvolumeinfos` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/subvolumeinfos"
```


### Client Initialization

```go
client := subvolumeinfos.NewSubvolumeInfosClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubvolumeInfosClient.SubvolumesCreate`

```go
ctx := context.TODO()
id := subvolumeinfos.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "subVolumeName")

payload := subvolumeinfos.SubvolumeInfo{
	// ...
}


if err := client.SubvolumesCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SubvolumeInfosClient.SubvolumesDelete`

```go
ctx := context.TODO()
id := subvolumeinfos.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "subVolumeName")

if err := client.SubvolumesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SubvolumeInfosClient.SubvolumesGet`

```go
ctx := context.TODO()
id := subvolumeinfos.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "subVolumeName")

read, err := client.SubvolumesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubvolumeInfosClient.SubvolumesGetMetadata`

```go
ctx := context.TODO()
id := subvolumeinfos.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "subVolumeName")

if err := client.SubvolumesGetMetadataThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SubvolumeInfosClient.SubvolumesListByVolume`

```go
ctx := context.TODO()
id := subvolumeinfos.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

// alternatively `client.SubvolumesListByVolume(ctx, id)` can be used to do batched pagination
items, err := client.SubvolumesListByVolumeComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SubvolumeInfosClient.SubvolumesUpdate`

```go
ctx := context.TODO()
id := subvolumeinfos.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName", "subVolumeName")

payload := subvolumeinfos.SubvolumePatchRequest{
	// ...
}


if err := client.SubvolumesUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
