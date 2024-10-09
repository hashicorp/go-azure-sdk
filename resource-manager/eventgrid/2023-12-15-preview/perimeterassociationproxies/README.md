
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/perimeterassociationproxies` Documentation

The `perimeterassociationproxies` SDK allows for interaction with Azure Resource Manager `eventgrid` (API Version `2023-12-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/perimeterassociationproxies"
```


### Client Initialization

```go
client := perimeterassociationproxies.NewPerimeterAssociationProxiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PerimeterAssociationProxiesClient.NetworkSecurityPerimeterConfigurationsGet`

```go
ctx := context.TODO()
id := perimeterassociationproxies.NewScopedNetworkSecurityPerimeterConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "networkSecurityPerimeterConfigurationName")

read, err := client.NetworkSecurityPerimeterConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PerimeterAssociationProxiesClient.NetworkSecurityPerimeterConfigurationsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.NetworkSecurityPerimeterConfigurationsList(ctx, id)` can be used to do batched pagination
items, err := client.NetworkSecurityPerimeterConfigurationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PerimeterAssociationProxiesClient.NetworkSecurityPerimeterConfigurationsReconcile`

```go
ctx := context.TODO()
id := perimeterassociationproxies.NewScopedNetworkSecurityPerimeterConfigurationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "networkSecurityPerimeterConfigurationName")

if err := client.NetworkSecurityPerimeterConfigurationsReconcileThenPoll(ctx, id); err != nil {
	// handle the error
}
```
