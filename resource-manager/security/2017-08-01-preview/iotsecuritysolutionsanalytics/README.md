
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


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue")

read, err := client.IoTSecuritySolutionsAnalyticsAggregatedAlertDismiss(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsAggregatedAlertGet`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedAlertID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedAlertValue")

read, err := client.IoTSecuritySolutionsAnalyticsAggregatedAlertGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsAggregatedAlertsList`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

// alternatively `client.IoTSecuritySolutionsAnalyticsAggregatedAlertsList(ctx, id, iotsecuritysolutionsanalytics.DefaultIoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions())` can be used to do batched pagination
items, err := client.IoTSecuritySolutionsAnalyticsAggregatedAlertsListComplete(ctx, id, iotsecuritysolutionsanalytics.DefaultIoTSecuritySolutionsAnalyticsAggregatedAlertsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsGetAll`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

read, err := client.IoTSecuritySolutionsAnalyticsGetAll(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsGetDefault`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

read, err := client.IoTSecuritySolutionsAnalyticsGetDefault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsRecommendationGet`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewAggregatedRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue", "aggregatedRecommendationValue")

read, err := client.IoTSecuritySolutionsAnalyticsRecommendationGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IoTSecuritySolutionsAnalyticsClient.IoTSecuritySolutionsAnalyticsRecommendationsList`

```go
ctx := context.TODO()
id := iotsecuritysolutionsanalytics.NewIotSecuritySolutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "iotSecuritySolutionValue")

// alternatively `client.IoTSecuritySolutionsAnalyticsRecommendationsList(ctx, id, iotsecuritysolutionsanalytics.DefaultIoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions())` can be used to do batched pagination
items, err := client.IoTSecuritySolutionsAnalyticsRecommendationsListComplete(ctx, id, iotsecuritysolutionsanalytics.DefaultIoTSecuritySolutionsAnalyticsRecommendationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
