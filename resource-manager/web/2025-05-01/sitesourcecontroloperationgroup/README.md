
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontroloperationgroup` Documentation

The `sitesourcecontroloperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontroloperationgroup"
```


### Client Initialization

```go
client := sitesourcecontroloperationgroup.NewSiteSourceControlOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteSourceControlOperationGroupClient.WebAppsCreateOrUpdateSourceControl`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sitesourcecontroloperationgroup.SiteSourceControl{
	// ...
}


if err := client.WebAppsCreateOrUpdateSourceControlThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SiteSourceControlOperationGroupClient.WebAppsDeleteSourceControl`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsDeleteSourceControl(ctx, id, sitesourcecontroloperationgroup.DefaultWebAppsDeleteSourceControlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteSourceControlOperationGroupClient.WebAppsGetSourceControl`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetSourceControl(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteSourceControlOperationGroupClient.WebAppsUpdateSourceControl`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := sitesourcecontroloperationgroup.SiteSourceControl{
	// ...
}


read, err := client.WebAppsUpdateSourceControl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
