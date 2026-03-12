
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendationrules` Documentation

The `recommendationrules` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/recommendationrules"
```


### Client Initialization

```go
client := recommendationrules.NewRecommendationRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationRulesClient.RecommendationsDisableRecommendationForHostingEnvironment`

```go
ctx := context.TODO()
id := recommendationrules.NewHostingEnvironmentRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "recommendationName")

read, err := client.RecommendationsDisableRecommendationForHostingEnvironment(ctx, id, recommendationrules.DefaultRecommendationsDisableRecommendationForHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationRulesClient.RecommendationsGetRuleDetailsByHostingEnvironment`

```go
ctx := context.TODO()
id := recommendationrules.NewHostingEnvironmentRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "recommendationName")

read, err := client.RecommendationsGetRuleDetailsByHostingEnvironment(ctx, id, recommendationrules.DefaultRecommendationsGetRuleDetailsByHostingEnvironmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
