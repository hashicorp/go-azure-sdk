
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetroutes` Documentation

The `vnetroutes` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetroutes"
```


### Client Initialization

```go
client := vnetroutes.NewVnetRoutesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetRoutesClient.AppServicePlansCreateOrUpdateVnetRoute`

```go
ctx := context.TODO()
id := vnetroutes.NewRouteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "routeName")

payload := vnetroutes.VnetRoute{
	// ...
}


read, err := client.AppServicePlansCreateOrUpdateVnetRoute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetRoutesClient.AppServicePlansDeleteVnetRoute`

```go
ctx := context.TODO()
id := vnetroutes.NewRouteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "routeName")

read, err := client.AppServicePlansDeleteVnetRoute(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetRoutesClient.AppServicePlansGetRouteForVnet`

```go
ctx := context.TODO()
id := vnetroutes.NewRouteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "routeName")

read, err := client.AppServicePlansGetRouteForVnet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetRoutesClient.AppServicePlansListRoutesForVnet`

```go
ctx := context.TODO()
id := vnetroutes.NewServerFarmVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName")

read, err := client.AppServicePlansListRoutesForVnet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetRoutesClient.AppServicePlansUpdateVnetRoute`

```go
ctx := context.TODO()
id := vnetroutes.NewRouteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "routeName")

payload := vnetroutes.VnetRoute{
	// ...
}


read, err := client.AppServicePlansUpdateVnetRoute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
