
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2s` Documentation

The `siteauthsettingsv2s` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/siteauthsettingsv2s"
```


### Client Initialization

```go
client := siteauthsettingsv2s.NewSiteAuthSettingsV2sClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteAuthSettingsV2sClient.WebAppsGetAuthSettingsV2`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetAuthSettingsV2(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteAuthSettingsV2sClient.WebAppsGetAuthSettingsV2WithoutSecrets`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetAuthSettingsV2WithoutSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteAuthSettingsV2sClient.WebAppsUpdateAuthSettingsV2`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := siteauthsettingsv2s.SiteAuthSettingsV2{
	// ...
}


read, err := client.WebAppsUpdateAuthSettingsV2(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
