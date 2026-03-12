
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessmoduleslotoperationgroup` Documentation

The `instanceprocessmoduleslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessmoduleslotoperationgroup"
```


### Client Initialization

```go
client := instanceprocessmoduleslotoperationgroup.NewInstanceProcessModuleSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstanceProcessModuleSlotOperationGroupClient.WebAppsGetInstanceProcessModuleSlot`

```go
ctx := context.TODO()
id := instanceprocessmoduleslotoperationgroup.NewSlotInstanceProcessModuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId", "moduleName")

read, err := client.WebAppsGetInstanceProcessModuleSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstanceProcessModuleSlotOperationGroupClient.WebAppsListInstanceProcessModulesSlot`

```go
ctx := context.TODO()
id := instanceprocessmoduleslotoperationgroup.NewSlotInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId")

// alternatively `client.WebAppsListInstanceProcessModulesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceProcessModulesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
