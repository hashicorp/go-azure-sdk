
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitecontaineroperationgroup` Documentation

The `sitecontaineroperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sitecontaineroperationgroup"
```


### Client Initialization

```go
client := sitecontaineroperationgroup.NewSiteContainerOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteContainerOperationGroupClient.WebAppsCreateOrUpdateSiteContainerSlot`

```go
ctx := context.TODO()
id := sitecontaineroperationgroup.NewSlotSitecontainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "sitecontainerName")

payload := sitecontaineroperationgroup.SiteContainer{
	// ...
}


read, err := client.WebAppsCreateOrUpdateSiteContainerSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContainerOperationGroupClient.WebAppsDeleteSiteContainerSlot`

```go
ctx := context.TODO()
id := sitecontaineroperationgroup.NewSlotSitecontainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "sitecontainerName")

read, err := client.WebAppsDeleteSiteContainerSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContainerOperationGroupClient.WebAppsGetSiteContainerSlot`

```go
ctx := context.TODO()
id := sitecontaineroperationgroup.NewSlotSitecontainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "sitecontainerName")

read, err := client.WebAppsGetSiteContainerSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteContainerOperationGroupClient.WebAppsListSiteContainersSlot`

```go
ctx := context.TODO()
id := sitecontaineroperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListSiteContainersSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListSiteContainersSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
