
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetgateways` Documentation

The `vnetgateways` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetgateways"
```


### Client Initialization

```go
client := vnetgateways.NewVnetGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetGatewaysClient.AppServicePlansGetVnetGateway`

```go
ctx := context.TODO()
id := vnetgateways.NewVirtualNetworkConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "gatewayName")

read, err := client.AppServicePlansGetVnetGateway(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetGatewaysClient.AppServicePlansUpdateVnetGateway`

```go
ctx := context.TODO()
id := vnetgateways.NewVirtualNetworkConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName", "gatewayName")

payload := vnetgateways.VnetGateway{
	// ...
}


read, err := client.AppServicePlansUpdateVnetGateway(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
