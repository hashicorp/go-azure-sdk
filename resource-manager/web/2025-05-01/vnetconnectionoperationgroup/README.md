
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetconnectionoperationgroup` Documentation

The `vnetconnectionoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetconnectionoperationgroup"
```


### Client Initialization

```go
client := vnetconnectionoperationgroup.NewVnetConnectionOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetConnectionOperationGroupClient.WebAppsCreateOrUpdateVnetConnection`

```go
ctx := context.TODO()
id := vnetconnectionoperationgroup.NewVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "virtualNetworkConnectionName")

payload := vnetconnectionoperationgroup.VnetInfoResource{
	// ...
}


read, err := client.WebAppsCreateOrUpdateVnetConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionOperationGroupClient.WebAppsDeleteVnetConnection`

```go
ctx := context.TODO()
id := vnetconnectionoperationgroup.NewVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "virtualNetworkConnectionName")

read, err := client.WebAppsDeleteVnetConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionOperationGroupClient.WebAppsGetVnetConnection`

```go
ctx := context.TODO()
id := vnetconnectionoperationgroup.NewVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "virtualNetworkConnectionName")

read, err := client.WebAppsGetVnetConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionOperationGroupClient.WebAppsListVnetConnections`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsListVnetConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetConnectionOperationGroupClient.WebAppsUpdateVnetConnection`

```go
ctx := context.TODO()
id := vnetconnectionoperationgroup.NewVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "virtualNetworkConnectionName")

payload := vnetconnectionoperationgroup.VnetInfoResource{
	// ...
}


read, err := client.WebAppsUpdateVnetConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
