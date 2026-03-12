
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2operationgroup` Documentation

The `siteauthsettingsv2operationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2operationgroup"
```


### Client Initialization

```go
client := siteauthsettingsv2operationgroup.NewSiteAuthSettingsV2OperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteAuthSettingsV2OperationGroupClient.WebAppsGetAuthSettingsV2Slot`

```go
ctx := context.TODO()
id := siteauthsettingsv2operationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetAuthSettingsV2Slot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteAuthSettingsV2OperationGroupClient.WebAppsGetAuthSettingsV2WithoutSecretsSlot`

```go
ctx := context.TODO()
id := siteauthsettingsv2operationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetAuthSettingsV2WithoutSecretsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteAuthSettingsV2OperationGroupClient.WebAppsUpdateAuthSettingsV2Slot`

```go
ctx := context.TODO()
id := siteauthsettingsv2operationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := siteauthsettingsv2operationgroup.SiteAuthSettingsV2{
	// ...
}


read, err := client.WebAppsUpdateAuthSettingsV2Slot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
