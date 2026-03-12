
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitelogsconfigs` Documentation

The `sitelogsconfigs` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitelogsconfigs"
```


### Client Initialization

```go
client := sitelogsconfigs.NewSiteLogsConfigsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteLogsConfigsClient.WebAppsGetDiagnosticLogsConfiguration`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetDiagnosticLogsConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteLogsConfigsClient.WebAppsUpdateDiagnosticLogsConfig`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sitelogsconfigs.SiteLogsConfig{
	// ...
}


read, err := client.WebAppsUpdateDiagnosticLogsConfig(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
