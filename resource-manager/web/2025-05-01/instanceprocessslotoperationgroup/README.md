
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessslotoperationgroup` Documentation

The `instanceprocessslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/instanceprocessslotoperationgroup"
```


### Client Initialization

```go
client := instanceprocessslotoperationgroup.NewInstanceProcessSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstanceProcessSlotOperationGroupClient.WebAppsDeleteInstanceProcessSlot`

```go
ctx := context.TODO()
id := instanceprocessslotoperationgroup.NewSlotInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId")

read, err := client.WebAppsDeleteInstanceProcessSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstanceProcessSlotOperationGroupClient.WebAppsGetInstanceProcessDumpSlot`

```go
ctx := context.TODO()
id := instanceprocessslotoperationgroup.NewSlotInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId")

read, err := client.WebAppsGetInstanceProcessDumpSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstanceProcessSlotOperationGroupClient.WebAppsGetInstanceProcessSlot`

```go
ctx := context.TODO()
id := instanceprocessslotoperationgroup.NewSlotInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId")

read, err := client.WebAppsGetInstanceProcessSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstanceProcessSlotOperationGroupClient.WebAppsListInstanceProcessThreadsSlot`

```go
ctx := context.TODO()
id := instanceprocessslotoperationgroup.NewSlotInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId", "processId")

// alternatively `client.WebAppsListInstanceProcessThreadsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceProcessThreadsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InstanceProcessSlotOperationGroupClient.WebAppsListInstanceProcessesSlot`

```go
ctx := context.TODO()
id := instanceprocessslotoperationgroup.NewSlotInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "instanceId")

// alternatively `client.WebAppsListInstanceProcessesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceProcessesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
