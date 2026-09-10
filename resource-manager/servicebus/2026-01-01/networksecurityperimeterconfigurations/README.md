
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/networksecurityperimeterconfigurations` Documentation

The `networksecurityperimeterconfigurations` SDK allows for interaction with Azure Resource Manager `servicebus` (API Version `2026-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicebus/2026-01-01/networksecurityperimeterconfigurations"
```


### Client Initialization

```go
client := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.GetResourceAssociationName`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "networkSecurityPerimeterConfigurationName")

read, err := client.GetResourceAssociationName(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterConfigurationList`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.NetworkSecurityPerimeterConfigurationList(ctx, id)` can be used to do batched pagination
items, err := client.NetworkSecurityPerimeterConfigurationListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.Reconcile`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "networkSecurityPerimeterConfigurationName")

read, err := client.Reconcile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
