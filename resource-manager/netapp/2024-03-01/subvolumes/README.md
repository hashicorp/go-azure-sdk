
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2024-03-01/subvolumes` Documentation

The `subvolumes` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2024-03-01/subvolumes"
```


### Client Initialization

```go
client := subvolumes.NewSubVolumesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubVolumesClient.Create`

```go
ctx := context.TODO()
id := subvolumes.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue", "subVolumeValue")

payload := subvolumes.SubvolumeInfo{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SubVolumesClient.Delete`

```go
ctx := context.TODO()
id := subvolumes.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue", "subVolumeValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SubVolumesClient.Get`

```go
ctx := context.TODO()
id := subvolumes.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue", "subVolumeValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubVolumesClient.GetMetadata`

```go
ctx := context.TODO()
id := subvolumes.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue", "subVolumeValue")

if err := client.GetMetadataThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SubVolumesClient.ListByVolume`

```go
ctx := context.TODO()
id := subvolumes.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue")

// alternatively `client.ListByVolume(ctx, id)` can be used to do batched pagination
items, err := client.ListByVolumeComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SubVolumesClient.Update`

```go
ctx := context.TODO()
id := subvolumes.NewSubVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue", "subVolumeValue")

payload := subvolumes.SubvolumePatchRequest{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
