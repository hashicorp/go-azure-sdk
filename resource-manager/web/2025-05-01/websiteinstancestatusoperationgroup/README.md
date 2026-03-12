
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/websiteinstancestatusoperationgroup` Documentation

The `websiteinstancestatusoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/websiteinstancestatusoperationgroup"
```


### Client Initialization

```go
client := websiteinstancestatusoperationgroup.NewWebSiteInstanceStatusOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WebSiteInstanceStatusOperationGroupClient.WebAppsGetInstanceInfoSlot`

```go
ctx := context.TODO()
id := websiteinstancestatusoperationgroup.NewSlotInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId")

read, err := client.WebAppsGetInstanceInfoSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WebSiteInstanceStatusOperationGroupClient.WebAppsListInstanceIdentifiersSlot`

```go
ctx := context.TODO()
id := websiteinstancestatusoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListInstanceIdentifiersSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceIdentifiersSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
