
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeter` Documentation

The `networksecurityperimeter` SDK allows for interaction with the Azure Resource Manager Service `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/networksecurityperimeter"
```


### Client Initialization

```go
client := networksecurityperimeter.NewNetworkSecurityPerimeterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkSecurityPerimeterClient.GetConfiguration`

```go
ctx := context.TODO()
id := networksecurityperimeter.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "networkSecurityPerimeterConfigurationValue")

read, err := client.GetConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkSecurityPerimeterClient.ListConfigurations`

```go
ctx := context.TODO()
id := networksecurityperimeter.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue")

// alternatively `client.ListConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.ListConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkSecurityPerimeterClient.ReconcileConfiguration`

```go
ctx := context.TODO()
id := networksecurityperimeter.NewNetworkSecurityPerimeterConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "networkSecurityPerimeterConfigurationValue")

if err := client.ReconcileConfigurationThenPoll(ctx, id); err != nil {
	// handle the error
}
```
