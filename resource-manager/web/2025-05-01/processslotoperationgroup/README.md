
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processslotoperationgroup` Documentation

The `processslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processslotoperationgroup"
```


### Client Initialization

```go
client := processslotoperationgroup.NewProcessSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProcessSlotOperationGroupClient.WebAppsDeleteProcessSlot`

```go
ctx := context.TODO()
id := processslotoperationgroup.NewSlotProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId")

read, err := client.WebAppsDeleteProcessSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessSlotOperationGroupClient.WebAppsGetProcessDumpSlot`

```go
ctx := context.TODO()
id := processslotoperationgroup.NewSlotProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId")

read, err := client.WebAppsGetProcessDumpSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessSlotOperationGroupClient.WebAppsGetProcessSlot`

```go
ctx := context.TODO()
id := processslotoperationgroup.NewSlotProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId")

read, err := client.WebAppsGetProcessSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessSlotOperationGroupClient.WebAppsListProcessThreadsSlot`

```go
ctx := context.TODO()
id := processslotoperationgroup.NewSlotProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "processId")

// alternatively `client.WebAppsListProcessThreadsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListProcessThreadsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProcessSlotOperationGroupClient.WebAppsListProcessesSlot`

```go
ctx := context.TODO()
id := processslotoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListProcessesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListProcessesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
