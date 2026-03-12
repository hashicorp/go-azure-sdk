
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontrols` Documentation

The `sitesourcecontrols` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitesourcecontrols"
```


### Client Initialization

```go
client := sitesourcecontrols.NewSiteSourceControlsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteSourceControlsClient.WebAppsCreateOrUpdateSourceControlSlot`

```go
ctx := context.TODO()
id := sitesourcecontrols.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := sitesourcecontrols.SiteSourceControl{
	// ...
}


if err := client.WebAppsCreateOrUpdateSourceControlSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SiteSourceControlsClient.WebAppsDeleteSourceControlSlot`

```go
ctx := context.TODO()
id := sitesourcecontrols.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsDeleteSourceControlSlot(ctx, id, sitesourcecontrols.DefaultWebAppsDeleteSourceControlSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteSourceControlsClient.WebAppsGetSourceControlSlot`

```go
ctx := context.TODO()
id := sitesourcecontrols.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetSourceControlSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteSourceControlsClient.WebAppsUpdateSourceControlSlot`

```go
ctx := context.TODO()
id := sitesourcecontrols.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := sitesourcecontrols.SiteSourceControl{
	// ...
}


read, err := client.WebAppsUpdateSourceControlSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
