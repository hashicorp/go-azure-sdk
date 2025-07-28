
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/loganalytics` Documentation

The `loganalytics` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/loganalytics"
```


### Client Initialization

```go
client := loganalytics.NewLogAnalyticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogAnalyticsClient.GetLogAnalyticsLocations`

```go
ctx := context.TODO()
id := loganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.GetLogAnalyticsLocations(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogAnalyticsClient.GetLogAnalyticsMetrics`

```go
ctx := context.TODO()
id := loganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.GetLogAnalyticsMetrics(ctx, id, loganalytics.DefaultGetLogAnalyticsMetricsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogAnalyticsClient.GetLogAnalyticsRankings`

```go
ctx := context.TODO()
id := loganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.GetLogAnalyticsRankings(ctx, id, loganalytics.DefaultGetLogAnalyticsRankingsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LogAnalyticsClient.GetLogAnalyticsResources`

```go
ctx := context.TODO()
id := loganalytics.NewProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName")

read, err := client.GetLogAnalyticsResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
