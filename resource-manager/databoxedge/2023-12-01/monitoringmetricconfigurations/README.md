
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/monitoringmetricconfigurations` Documentation

The `monitoringmetricconfigurations` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/monitoringmetricconfigurations"
```


### Client Initialization

```go
client := monitoringmetricconfigurations.NewMonitoringMetricConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitoringMetricConfigurationsClient.MonitoringConfigCreateOrUpdate`

```go
ctx := context.TODO()
id := monitoringmetricconfigurations.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "roleName")

payload := monitoringmetricconfigurations.MonitoringMetricConfiguration{
	// ...
}


if err := client.MonitoringConfigCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoringMetricConfigurationsClient.MonitoringConfigDelete`

```go
ctx := context.TODO()
id := monitoringmetricconfigurations.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "roleName")

if err := client.MonitoringConfigDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoringMetricConfigurationsClient.MonitoringConfigGet`

```go
ctx := context.TODO()
id := monitoringmetricconfigurations.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "roleName")

read, err := client.MonitoringConfigGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringMetricConfigurationsClient.MonitoringConfigList`

```go
ctx := context.TODO()
id := monitoringmetricconfigurations.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "roleName")

// alternatively `client.MonitoringConfigList(ctx, id)` can be used to do batched pagination
items, err := client.MonitoringConfigListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
