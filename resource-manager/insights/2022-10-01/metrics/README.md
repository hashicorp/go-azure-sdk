
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/metrics` Documentation

The `metrics` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2022-10-01/metrics"
```


### Client Initialization

```go
client := metrics.NewMetricsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MetricsClient.PredictiveMetricGet`

```go
ctx := context.TODO()
id := metrics.NewAutoScaleSettingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "autoScaleSettingName")

read, err := client.PredictiveMetricGet(ctx, id, metrics.DefaultPredictiveMetricGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
