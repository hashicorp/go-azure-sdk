
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/bandwidthschedules` Documentation

The `bandwidthschedules` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/bandwidthschedules"
```


### Client Initialization

```go
client := bandwidthschedules.NewBandwidthSchedulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BandwidthSchedulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := bandwidthschedules.NewBandwidthScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "bandwidthScheduleValue")

payload := bandwidthschedules.BandwidthSchedule{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BandwidthSchedulesClient.Delete`

```go
ctx := context.TODO()
id := bandwidthschedules.NewBandwidthScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "bandwidthScheduleValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BandwidthSchedulesClient.Get`

```go
ctx := context.TODO()
id := bandwidthschedules.NewBandwidthScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "bandwidthScheduleValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BandwidthSchedulesClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := bandwidthschedules.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
