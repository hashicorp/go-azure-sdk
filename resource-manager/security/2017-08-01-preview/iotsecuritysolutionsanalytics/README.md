
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutionsanalytics` Documentation

The `iotsecuritysolutionsanalytics` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutionsanalytics"
```


### Client Initialization

```go
client := iotsecuritysolutionsanalytics.NewIoTSecuritySolutionsAnalyticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.AggregatedAlertDismiss`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue")

read, err := client.AggregatedAlertDismiss(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.AggregatedAlertGet`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue")

read, err := client.AggregatedAlertGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.AggregatedAlertsList`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

// alternatively `client.AggregatedAlertsList(ctx, id, iotsecuritysolutionsanalytics.DefaultAggregatedAlertsListOperationOptions())` can be used to do batched pagination
items, err := client.AggregatedAlertsListComplete(ctx, id, iotsecuritysolutionsanalytics.DefaultAggregatedAlertsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.GetAll`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

// alternatively `client.GetAll(ctx, id)` can be used to do batched pagination
items, err := client.GetAllComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.GetDefault`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

read, err := client.GetDefault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.RecommendationGet`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedRecommendationValue")

read, err := client.RecommendationGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.RecommendationsList`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

// alternatively `client.RecommendationsList(ctx, id, iotsecuritysolutionsanalytics.DefaultRecommendationsListOperationOptions())` can be used to do batched pagination
items, err := client.RecommendationsListComplete(ctx, id, iotsecuritysolutionsanalytics.DefaultRecommendationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
