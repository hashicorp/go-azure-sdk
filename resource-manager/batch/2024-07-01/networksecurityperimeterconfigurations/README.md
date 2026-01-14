
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeterconfigurations` Documentation

The `networksecurityperimeterconfigurations` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeterconfigurations"
```


### Client Initialization

```go
client := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterGetConfiguration`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "networkSecurityPerimeterConfigurationName")

read, err := client.NetworkSecurityPerimeterGetConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterListConfigurations`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName")

// alternatively `client.NetworkSecurityPerimeterListConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.NetworkSecurityPerimeterListConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkSecurityPerimeterConfigurationsClient.NetworkSecurityPerimeterReconcileConfiguration`

```go
ctx := context.TODO()
id := networksecurityperimeterconfigurations.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "networkSecurityPerimeterConfigurationName")

if err := client.NetworkSecurityPerimeterReconcileConfigurationThenPoll(ctx, id); err != nil {
	// handle the error
}
```
