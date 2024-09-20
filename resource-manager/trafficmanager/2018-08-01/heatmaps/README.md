
## `github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/heatmaps` Documentation

The `heatmaps` SDK allows for interaction with Azure Resource Manager `trafficmanager` (API Version `2018-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/trafficmanager/2018-08-01/heatmaps"
```


### Client Initialization

```go
client := heatmaps.NewHeatMapsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HeatMapsClient.HeatMapGet`

```go
ctx := context.TODO()
id := heatmaps.NewTrafficManagerProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.HeatMapGet(ctx, id, heatmaps.DefaultHeatMapGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
