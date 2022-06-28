
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/flux` Documentation

The `flux` SDK allows for interaction with the Azure Resource Manager Service `kubernetesconfiguration` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/flux"
```


### Client Initialization

```go
client := flux.NewFluxClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `FluxClient.ConfigurationsCreateOrUpdate`

```go
ctx := context.TODO()
id := flux.NewFluxConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "fluxConfigurationValue")

payload := flux.FluxConfiguration{
	// ...
}

future, err := client.ConfigurationsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FluxClient.ConfigurationsDelete`

```go
ctx := context.TODO()
id := flux.NewFluxConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "fluxConfigurationValue")
future, err := client.ConfigurationsDelete(ctx, id, flux.DefaultConfigurationsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FluxClient.ConfigurationsGet`

```go
ctx := context.TODO()
id := flux.NewFluxConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "fluxConfigurationValue")
read, err := client.ConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FluxClient.ConfigurationsList`

```go
ctx := context.TODO()
id := flux.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue")
// alternatively `client.ConfigurationsList(ctx, id)` can be used to do batched pagination
items, err := client.ConfigurationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FluxClient.ConfigurationsUpdate`

```go
ctx := context.TODO()
id := flux.NewFluxConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "fluxConfigurationValue")

payload := flux.FluxConfigurationPatch{
	// ...
}

future, err := client.ConfigurationsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
