
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermetricsconfigurations` Documentation

The `clustermetricsconfigurations` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/clustermetricsconfigurations"
```


### Client Initialization

```go
client := clustermetricsconfigurations.NewClusterMetricsConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterMetricsConfigurationsClient.MetricsConfigurationsCreateOrUpdate`

```go
ctx := context.TODO()
id := clustermetricsconfigurations.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

payload := clustermetricsconfigurations.ClusterMetricsConfiguration{
	// ...
}


if err := client.MetricsConfigurationsCreateOrUpdateThenPoll(ctx, id, payload, clustermetricsconfigurations.DefaultMetricsConfigurationsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterMetricsConfigurationsClient.MetricsConfigurationsDelete`

```go
ctx := context.TODO()
id := clustermetricsconfigurations.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

if err := client.MetricsConfigurationsDeleteThenPoll(ctx, id, clustermetricsconfigurations.DefaultMetricsConfigurationsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterMetricsConfigurationsClient.MetricsConfigurationsGet`

```go
ctx := context.TODO()
id := clustermetricsconfigurations.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

read, err := client.MetricsConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterMetricsConfigurationsClient.MetricsConfigurationsListByCluster`

```go
ctx := context.TODO()
id := clustermetricsconfigurations.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.MetricsConfigurationsListByCluster(ctx, id, clustermetricsconfigurations.DefaultMetricsConfigurationsListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.MetricsConfigurationsListByClusterComplete(ctx, id, clustermetricsconfigurations.DefaultMetricsConfigurationsListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterMetricsConfigurationsClient.MetricsConfigurationsUpdate`

```go
ctx := context.TODO()
id := clustermetricsconfigurations.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

payload := clustermetricsconfigurations.ClusterMetricsConfigurationPatchParameters{
	// ...
}


if err := client.MetricsConfigurationsUpdateThenPoll(ctx, id, payload, clustermetricsconfigurations.DefaultMetricsConfigurationsUpdateOperationOptions()); err != nil {
	// handle the error
}
```
