
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigsnapshotslotresourceoperationgroup` Documentation

The `siteconfigsnapshotslotresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigsnapshotslotresourceoperationgroup"
```


### Client Initialization

```go
client := siteconfigsnapshotslotresourceoperationgroup.NewSiteConfigSnapshotSlotResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteConfigSnapshotSlotResourceOperationGroupClient.WebAppsGetConfigurationSnapshotSlot`

```go
ctx := context.TODO()
id := siteconfigsnapshotslotresourceoperationgroup.NewWebSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "snapshotId")

read, err := client.WebAppsGetConfigurationSnapshotSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConfigSnapshotSlotResourceOperationGroupClient.WebAppsListConfigurationsSlot`

```go
ctx := context.TODO()
id := siteconfigsnapshotslotresourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListConfigurationsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListConfigurationsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteConfigSnapshotSlotResourceOperationGroupClient.WebAppsRecoverSiteConfigurationSnapshotSlot`

```go
ctx := context.TODO()
id := siteconfigsnapshotslotresourceoperationgroup.NewWebSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "snapshotId")

read, err := client.WebAppsRecoverSiteConfigurationSnapshotSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
