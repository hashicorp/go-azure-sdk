
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworkgateways` Documentation

The `workloadnetworkgateways` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworkgateways"
```


### Client Initialization

```go
client := workloadnetworkgateways.NewWorkloadNetworkGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadNetworkGatewaysClient.WorkloadNetworksGetGateway`

```go
ctx := context.TODO()
id := workloadnetworkgateways.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "gatewayId")

read, err := client.WorkloadNetworksGetGateway(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworkGatewaysClient.WorkloadNetworksListGateways`

```go
ctx := context.TODO()
id := workloadnetworkgateways.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.WorkloadNetworksListGateways(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListGatewaysComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
