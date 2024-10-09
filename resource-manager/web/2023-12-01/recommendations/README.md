
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/recommendations` Documentation

The `recommendations` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/recommendations"
```


### Client Initialization

```go
client := recommendations.NewRecommendationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationsClient.DisableAllForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.DisableAllForHostingEnvironment(ctx, id, recommendations.DefaultDisableAllForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.DisableAllForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.DisableAllForWebApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.DisableRecommendationForHostingEnvironment`

```go
ctx := context.TODO()
id := recommendations.NewHostingEnvironmentRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "recommendationName")

read, err := client.DisableRecommendationForHostingEnvironment(ctx, id, recommendations.DefaultDisableRecommendationForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.DisableRecommendationForSite`

```go
ctx := context.TODO()
id := recommendations.NewSiteRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "recommendationName")

read, err := client.DisableRecommendationForSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.DisableRecommendationForSubscription`

```go
ctx := context.TODO()
id := recommendations.NewRecommendationID("12345678-1234-9876-4563-123456789012", "recommendationName")

read, err := client.DisableRecommendationForSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.GetRuleDetailsByHostingEnvironment`

```go
ctx := context.TODO()
id := recommendations.NewHostingEnvironmentRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "recommendationName")

read, err := client.GetRuleDetailsByHostingEnvironment(ctx, id, recommendations.DefaultGetRuleDetailsByHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.GetRuleDetailsByWebApp`

```go
ctx := context.TODO()
id := recommendations.NewSiteRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "recommendationName")

read, err := client.GetRuleDetailsByWebApp(ctx, id, recommendations.DefaultGetRuleDetailsByWebAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, recommendations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, recommendations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationsClient.ListHistoryForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.ListHistoryForHostingEnvironment(ctx, id, recommendations.DefaultListHistoryForHostingEnvironmentOperationOptions())` can be used to do batched pagination
items, err := client.ListHistoryForHostingEnvironmentComplete(ctx, id, recommendations.DefaultListHistoryForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationsClient.ListHistoryForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.ListHistoryForWebApp(ctx, id, recommendations.DefaultListHistoryForWebAppOperationOptions())` can be used to do batched pagination
items, err := client.ListHistoryForWebAppComplete(ctx, id, recommendations.DefaultListHistoryForWebAppOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationsClient.ListRecommendedRulesForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.ListRecommendedRulesForHostingEnvironment(ctx, id, recommendations.DefaultListRecommendedRulesForHostingEnvironmentOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendedRulesForHostingEnvironmentComplete(ctx, id, recommendations.DefaultListRecommendedRulesForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationsClient.ListRecommendedRulesForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.ListRecommendedRulesForWebApp(ctx, id, recommendations.DefaultListRecommendedRulesForWebAppOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendedRulesForWebAppComplete(ctx, id, recommendations.DefaultListRecommendedRulesForWebAppOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationsClient.ResetAllFilters`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ResetAllFilters(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.ResetAllFiltersForHostingEnvironment`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

read, err := client.ResetAllFiltersForHostingEnvironment(ctx, id, recommendations.DefaultResetAllFiltersForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.ResetAllFiltersForWebApp`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.ResetAllFiltersForWebApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
