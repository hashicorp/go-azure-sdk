
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/resourcehealthmetadata` Documentation

The `resourcehealthmetadata` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/resourcehealthmetadata"
```


### Client Initialization

```go
client := resourcehealthmetadata.NewResourceHealthMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceHealthMetadataClient.GetBySite`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue")

read, err := client.GetBySite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceHealthMetadataClient.GetBySiteSlot`

```go
ctx := context.TODO()
id := resourcehealthmetadata.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue")

read, err := client.GetBySiteSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceHealthMetadataClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceHealthMetadataClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceHealthMetadataClient.ListBySite`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue")

// alternatively `client.ListBySite(ctx, id)` can be used to do batched pagination
items, err := client.ListBySiteComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceHealthMetadataClient.ListBySiteSlot`

```go
ctx := context.TODO()
id := resourcehealthmetadata.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue")

// alternatively `client.ListBySiteSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListBySiteSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
