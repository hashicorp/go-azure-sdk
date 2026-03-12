
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworkoperationgroup` Documentation

The `swiftvirtualnetworkoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworkoperationgroup"
```


### Client Initialization

```go
client := swiftvirtualnetworkoperationgroup.NewSwiftVirtualNetworkOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SwiftVirtualNetworkOperationGroupClient.WebAppsCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot`

```go
ctx := context.TODO()
id := swiftvirtualnetworkoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := swiftvirtualnetworkoperationgroup.SwiftVirtualNetwork{
	// ...
}


read, err := client.WebAppsCreateOrUpdateSwiftVirtualNetworkConnectionWithCheckSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworkOperationGroupClient.WebAppsDeleteSwiftVirtualNetworkSlot`

```go
ctx := context.TODO()
id := swiftvirtualnetworkoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsDeleteSwiftVirtualNetworkSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworkOperationGroupClient.WebAppsGetSwiftVirtualNetworkConnectionSlot`

```go
ctx := context.TODO()
id := swiftvirtualnetworkoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsGetSwiftVirtualNetworkConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworkOperationGroupClient.WebAppsUpdateSwiftVirtualNetworkConnectionWithCheckSlot`

```go
ctx := context.TODO()
id := swiftvirtualnetworkoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

payload := swiftvirtualnetworkoperationgroup.SwiftVirtualNetwork{
	// ...
}


read, err := client.WebAppsUpdateSwiftVirtualNetworkConnectionWithCheckSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
