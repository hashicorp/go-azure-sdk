
## `github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-08-15/outboundnetworkdependenciesendpoints` Documentation

The `outboundnetworkdependenciesendpoints` SDK allows for interaction with the Azure Resource Manager Service `kusto` (API Version `2023-08-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-08-15/outboundnetworkdependenciesendpoints"
```


### Client Initialization

```go
client := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OutboundNetworkDependenciesEndpointsClient.ClustersListOutboundNetworkDependenciesEndpoints`

```go
ctx := context.TODO()
id := commonids.NewKustoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

// alternatively `client.ClustersListOutboundNetworkDependenciesEndpoints(ctx, id)` can be used to do batched pagination
items, err := client.ClustersListOutboundNetworkDependenciesEndpointsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
