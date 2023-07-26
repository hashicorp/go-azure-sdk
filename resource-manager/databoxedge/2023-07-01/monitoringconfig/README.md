
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/monitoringconfig` Documentation

The `monitoringconfig` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/monitoringconfig"
```


### Client Initialization

```go
client := monitoringconfig.NewMonitoringConfigClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitoringConfigClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := monitoringconfig.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "roleValue")

payload := monitoringconfig.MonitoringMetricConfiguration{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoringConfigClient.Delete`

```go
ctx := context.TODO()
id := monitoringconfig.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "roleValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MonitoringConfigClient.Get`

```go
ctx := context.TODO()
id := monitoringconfig.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "roleValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringConfigClient.List`

```go
ctx := context.TODO()
id := monitoringconfig.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "roleValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
