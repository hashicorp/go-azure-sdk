
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigslotresourceoperationgroup` Documentation

The `siteconfigslotresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigslotresourceoperationgroup"
```


### Client Initialization

```go
client := siteconfigslotresourceoperationgroup.NewSiteConfigSlotResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteConfigSlotResourceOperationGroupClient.WebAppsCreateOrUpdateConfigurationSlot`

```go
ctx := context.TODO()
id := siteconfigslotresourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := siteconfigslotresourceoperationgroup.SiteConfigResource{
	// ...
}


read, err := client.WebAppsCreateOrUpdateConfigurationSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConfigSlotResourceOperationGroupClient.WebAppsGetConfigurationSlot`

```go
ctx := context.TODO()
id := siteconfigslotresourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetConfigurationSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConfigSlotResourceOperationGroupClient.WebAppsListConfigurationSnapshotInfoSlot`

```go
ctx := context.TODO()
id := siteconfigslotresourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListConfigurationSnapshotInfoSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListConfigurationSnapshotInfoSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteConfigSlotResourceOperationGroupClient.WebAppsUpdateConfigurationSlot`

```go
ctx := context.TODO()
id := siteconfigslotresourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := siteconfigslotresourceoperationgroup.SiteConfigResource{
	// ...
}


read, err := client.WebAppsUpdateConfigurationSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
