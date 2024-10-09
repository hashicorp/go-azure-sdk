
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/alerts` Documentation

The `alerts` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/alerts"
```


### Client Initialization

```go
client := alerts.NewAlertsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AlertsClient.Get`

```go
ctx := context.TODO()
id := alerts.NewAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName", "alertName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AlertsClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := alerts.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceName")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
