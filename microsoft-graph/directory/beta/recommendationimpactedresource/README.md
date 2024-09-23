
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource` Documentation

The `recommendationimpactedresource` SDK allows for interaction with Microsoft Graph `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource"
```


### Client Initialization

```go
client := recommendationimpactedresource.NewRecommendationImpactedResourceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationId")

payload := recommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.CreateRecommendationImpactedResource(ctx, id, payload, recommendationimpactedresource.DefaultCreateRecommendationImpactedResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResourceComplete`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

read, err := client.CreateRecommendationImpactedResourceComplete(ctx, id, recommendationimpactedresource.DefaultCreateRecommendationImpactedResourceCompleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResourcePostpone`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

payload := recommendationimpactedresource.CreateRecommendationImpactedResourcePostponeRequest{
	// ...
}


read, err := client.CreateRecommendationImpactedResourcePostpone(ctx, id, payload, recommendationimpactedresource.DefaultCreateRecommendationImpactedResourcePostponeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResourceReactivate`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

read, err := client.CreateRecommendationImpactedResourceReactivate(ctx, id, recommendationimpactedresource.DefaultCreateRecommendationImpactedResourceReactivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.DeleteRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

read, err := client.DeleteRecommendationImpactedResource(ctx, id, recommendationimpactedresource.DefaultDeleteRecommendationImpactedResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.DismissRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

payload := recommendationimpactedresource.DismissRecommendationImpactedResourceRequest{
	// ...
}


read, err := client.DismissRecommendationImpactedResource(ctx, id, payload, recommendationimpactedresource.DefaultDismissRecommendationImpactedResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.GetRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

read, err := client.GetRecommendationImpactedResource(ctx, id, recommendationimpactedresource.DefaultGetRecommendationImpactedResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.GetRecommendationImpactedResourcesCount`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationId")

read, err := client.GetRecommendationImpactedResourcesCount(ctx, id, recommendationimpactedresource.DefaultGetRecommendationImpactedResourcesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.ListRecommendationImpactedResources`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationId")

// alternatively `client.ListRecommendationImpactedResources(ctx, id, recommendationimpactedresource.DefaultListRecommendationImpactedResourcesOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendationImpactedResourcesComplete(ctx, id, recommendationimpactedresource.DefaultListRecommendationImpactedResourcesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationImpactedResourceClient.UpdateRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationId", "impactedResourceId")

payload := recommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.UpdateRecommendationImpactedResource(ctx, id, payload, recommendationimpactedresource.DefaultUpdateRecommendationImpactedResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
