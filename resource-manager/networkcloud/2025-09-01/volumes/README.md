
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/volumes` Documentation

The `volumes` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/volumes"
```


### Client Initialization

```go
client := volumes.NewVolumesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VolumesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := volumes.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

payload := volumes.Volume{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, volumes.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesClient.Delete`

```go
ctx := context.TODO()
id := volumes.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

if err := client.DeleteThenPoll(ctx, id, volumes.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VolumesClient.Get`

```go
ctx := context.TODO()
id := volumes.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VolumesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, volumes.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, volumes.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VolumesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, volumes.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, volumes.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VolumesClient.Update`

```go
ctx := context.TODO()
id := volumes.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

payload := volumes.VolumePatchParameters{
	// ...
}


read, err := client.Update(ctx, id, payload, volumes.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
