
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfooperationgroup` Documentation

The `siteextensioninfooperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteextensioninfooperationgroup"
```


### Client Initialization

```go
client := siteextensioninfooperationgroup.NewSiteExtensionInfoOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteExtensionInfoOperationGroupClient.WebAppsDeleteSiteExtensionSlot`

```go
ctx := context.TODO()
id := siteextensioninfooperationgroup.NewSlotSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "siteExtensionId")

read, err := client.WebAppsDeleteSiteExtensionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteExtensionInfoOperationGroupClient.WebAppsGetSiteExtensionSlot`

```go
ctx := context.TODO()
id := siteextensioninfooperationgroup.NewSlotSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "siteExtensionId")

read, err := client.WebAppsGetSiteExtensionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteExtensionInfoOperationGroupClient.WebAppsInstallSiteExtensionSlot`

```go
ctx := context.TODO()
id := siteextensioninfooperationgroup.NewSlotSiteExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "siteExtensionId")

if err := client.WebAppsInstallSiteExtensionSlotThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SiteExtensionInfoOperationGroupClient.WebAppsListSiteExtensionsSlot`

```go
ctx := context.TODO()
id := siteextensioninfooperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListSiteExtensionsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSiteExtensionsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
