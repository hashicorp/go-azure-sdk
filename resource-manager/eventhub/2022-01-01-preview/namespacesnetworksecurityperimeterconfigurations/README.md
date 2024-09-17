
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespacesnetworksecurityperimeterconfigurations` Documentation

The `namespacesnetworksecurityperimeterconfigurations` SDK allows for interaction with Azure Resource Manager `eventhub` (API Version `2022-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/namespacesnetworksecurityperimeterconfigurations"
```


### Client Initialization

```go
client := namespacesnetworksecurityperimeterconfigurations.NewNamespacesNetworkSecurityPerimeterConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NamespacesNetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterConfigurationList`

```go
ctx := context.TODO()
id := namespacesnetworksecurityperimeterconfigurations.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.NetworkSecurityPerimeterConfigurationList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NamespacesNetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterConfigurationsCreateOrUpdate`

```go
ctx := context.TODO()
id := namespacesnetworksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "networkSecurityPerimeterConfigurationValue")

if err := client.NetworkSecurityPerimeterConfigurationsCreateOrUpdateThenPoll(ctx, id); err != nil {
	// handle the error
}
```
