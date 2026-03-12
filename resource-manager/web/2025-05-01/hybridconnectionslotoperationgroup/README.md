
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionslotoperationgroup` Documentation

The `hybridconnectionslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hybridconnectionslotoperationgroup"
```


### Client Initialization

```go
client := hybridconnectionslotoperationgroup.NewHybridConnectionSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HybridConnectionSlotOperationGroupClient.WebAppsCreateOrUpdateHybridConnectionSlot`

```go
ctx := context.TODO()
id := hybridconnectionslotoperationgroup.NewSlotHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionNamespaceName", "relayName")

payload := hybridconnectionslotoperationgroup.HybridConnection{
	// ...
}


read, err := client.WebAppsCreateOrUpdateHybridConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionSlotOperationGroupClient.WebAppsDeleteHybridConnectionSlot`

```go
ctx := context.TODO()
id := hybridconnectionslotoperationgroup.NewSlotHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionNamespaceName", "relayName")

read, err := client.WebAppsDeleteHybridConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionSlotOperationGroupClient.WebAppsGetHybridConnectionSlot`

```go
ctx := context.TODO()
id := hybridconnectionslotoperationgroup.NewSlotHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionNamespaceName", "relayName")

read, err := client.WebAppsGetHybridConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HybridConnectionSlotOperationGroupClient.WebAppsUpdateHybridConnectionSlot`

```go
ctx := context.TODO()
id := hybridconnectionslotoperationgroup.NewSlotHybridConnectionNamespaceRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hybridConnectionNamespaceName", "relayName")

payload := hybridconnectionslotoperationgroup.HybridConnection{
	// ...
}


read, err := client.WebAppsUpdateHybridConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
