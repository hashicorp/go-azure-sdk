
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesources` Documentation

The `vnetinforesources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/vnetinforesources"
```


### Client Initialization

```go
client := vnetinforesources.NewVnetInfoResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VnetInfoResourcesClient.AppServicePlansGetVnetFromServerFarm`

```go
ctx := context.TODO()
id := vnetinforesources.NewServerFarmVirtualNetworkConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName", "virtualNetworkConnectionName")

read, err := client.AppServicePlansGetVnetFromServerFarm(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VnetInfoResourcesClient.AppServicePlansListVnets`

```go
ctx := context.TODO()
id := commonids.NewAppServicePlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverFarmName")

read, err := client.AppServicePlansListVnets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
