
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesourceoperationgroup` Documentation

The `vnetinforesourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesourceoperationgroup"
```


### Client Initialization

```go
client := vnetinforesourceoperationgroup.NewVnetInfoResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetInfoResourceOperationGroupClient.WebAppsCreateOrUpdateVnetConnectionSlot`

```go
ctx := context.TODO()
id := vnetinforesourceoperationgroup.NewSlotVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "virtualNetworkConnectionName")

payload := vnetinforesourceoperationgroup.VnetInfoResource{
	// ...
}


read, err := client.WebAppsCreateOrUpdateVnetConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetInfoResourceOperationGroupClient.WebAppsDeleteVnetConnectionSlot`

```go
ctx := context.TODO()
id := vnetinforesourceoperationgroup.NewSlotVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "virtualNetworkConnectionName")

read, err := client.WebAppsDeleteVnetConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetInfoResourceOperationGroupClient.WebAppsGetVnetConnectionSlot`

```go
ctx := context.TODO()
id := vnetinforesourceoperationgroup.NewSlotVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "virtualNetworkConnectionName")

read, err := client.WebAppsGetVnetConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetInfoResourceOperationGroupClient.WebAppsListVnetConnectionsSlot`

```go
ctx := context.TODO()
id := vnetinforesourceoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.WebAppsListVnetConnectionsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetInfoResourceOperationGroupClient.WebAppsUpdateVnetConnectionSlot`

```go
ctx := context.TODO()
id := vnetinforesourceoperationgroup.NewSlotVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "virtualNetworkConnectionName")

payload := vnetinforesourceoperationgroup.VnetInfoResource{
	// ...
}


read, err := client.WebAppsUpdateVnetConnectionSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
