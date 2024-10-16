
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-09-01/wafloganalytics` Documentation

The `wafloganalytics` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-09-01/wafloganalytics"
```


### Client Initialization

```go
client := wafloganalytics.NewWafLogAnalyticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WafLogAnalyticsClient.LogAnalyticsGetWafLogAnalyticsMetrics`

```go
ctx := context.TODO()
id := wafloganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.LogAnalyticsGetWafLogAnalyticsMetrics(ctx, id, wafloganalytics.DefaultLogAnalyticsGetWafLogAnalyticsMetricsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WafLogAnalyticsClient.LogAnalyticsGetWafLogAnalyticsRankings`

```go
ctx := context.TODO()
id := wafloganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.LogAnalyticsGetWafLogAnalyticsRankings(ctx, id, wafloganalytics.DefaultLogAnalyticsGetWafLogAnalyticsRankingsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
