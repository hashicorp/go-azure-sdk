
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigresourceoperationgroup` Documentation

The `siteconfigresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteconfigresourceoperationgroup"
```


### Client Initialization

```go
client := siteconfigresourceoperationgroup.NewSiteConfigResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteConfigResourceOperationGroupClient.WebAppsGetConfigurationSnapshot`

```go
ctx := context.TODO()
id := siteconfigresourceoperationgroup.NewSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "snapshotId")

read, err := client.WebAppsGetConfigurationSnapshot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteConfigResourceOperationGroupClient.WebAppsListConfigurations`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteConfigResourceOperationGroupClient.WebAppsRecoverSiteConfigurationSnapshot`

```go
ctx := context.TODO()
id := siteconfigresourceoperationgroup.NewSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "snapshotId")

read, err := client.WebAppsRecoverSiteConfigurationSnapshot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
