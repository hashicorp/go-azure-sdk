
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendations` Documentation

The `recommendations` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendations"
```


### Client Initialization

```go
client := recommendations.NewRecommendationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
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
