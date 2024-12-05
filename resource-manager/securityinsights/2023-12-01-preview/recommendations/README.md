
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/recommendations` Documentation

The `recommendations` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/recommendations"
```


### Client Initialization

```go
client := recommendations.NewRecommendationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationsClient.GetRecommendationsList`

```go
ctx := context.TODO()
id := recommendations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.GetRecommendationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.GetSingleRecommendation`

```go
ctx := context.TODO()
id := recommendations.NewRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "recommendationId")

read, err := client.GetSingleRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationsClient.UpdateRecommendation`

```go
ctx := context.TODO()
id := recommendations.NewRecommendationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "recommendationId")
var payload []RecommendationPatch

if err := client.UpdateRecommendationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
