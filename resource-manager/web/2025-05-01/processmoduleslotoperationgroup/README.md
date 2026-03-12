
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleslotoperationgroup` Documentation

The `processmoduleslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleslotoperationgroup"
```


### Client Initialization

```go
client := processmoduleslotoperationgroup.NewProcessModuleSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProcessModuleSlotOperationGroupClient.WebAppsGetProcessModuleSlot`

```go
ctx := context.TODO()
id := processmoduleslotoperationgroup.NewProcessModuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId", "moduleName")

read, err := client.WebAppsGetProcessModuleSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessModuleSlotOperationGroupClient.WebAppsListProcessModulesSlot`

```go
ctx := context.TODO()
id := processmoduleslotoperationgroup.NewSlotProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId")

// alternatively `client.WebAppsListProcessModulesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListProcessModulesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
