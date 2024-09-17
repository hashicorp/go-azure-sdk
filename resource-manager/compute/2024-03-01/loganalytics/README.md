
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/loganalytics` Documentation

The `loganalytics` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-03-01/loganalytics"
```


### Client Initialization

```go
client := loganalytics.NewLogAnalyticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogAnalyticsClient.ExportRequestRateByInterval`

```go
ctx := context.TODO()
id := loganalytics.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := loganalytics.RequestRateByIntervalInput{
	// ...
}


if err := client.ExportRequestRateByIntervalThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LogAnalyticsClient.ExportThrottledRequests`

```go
ctx := context.TODO()
id := loganalytics.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := loganalytics.LogAnalyticsInputBase{
	// ...
}


if err := client.ExportThrottledRequestsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
