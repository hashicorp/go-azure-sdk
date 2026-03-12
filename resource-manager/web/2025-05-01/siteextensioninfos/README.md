
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfos` Documentation

The `siteextensioninfos` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfos"
```


### Client Initialization

```go
client := siteextensioninfos.NewSiteExtensionInfosClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteExtensionInfosClient.WebAppsDeleteSiteExtension`

```go
ctx := context.TODO()
id := siteextensioninfos.NewSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "siteExtensionId")

read, err := client.WebAppsDeleteSiteExtension(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteExtensionInfosClient.WebAppsGetSiteExtension`

```go
ctx := context.TODO()
id := siteextensioninfos.NewSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "siteExtensionId")

read, err := client.WebAppsGetSiteExtension(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteExtensionInfosClient.WebAppsInstallSiteExtension`

```go
ctx := context.TODO()
id := siteextensioninfos.NewSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "siteExtensionId")

if err := client.WebAppsInstallSiteExtensionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SiteExtensionInfosClient.WebAppsListSiteExtensions`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListSiteExtensions(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSiteExtensionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
