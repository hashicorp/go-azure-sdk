
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource` Documentation

The `recommendationimpactedresource` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendationimpactedresource"
```


### Client Initialization

```go
client := recommendationimpactedresource.NewRecommendationImpactedResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResource`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

payload := recommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.CreateRecommendationImpactedResource(ctx, id, payload)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.CreateRecommendationImpactedResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.CreateRecommendationImpactedResourceDismis`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := recommendationimpactedresource.CreateRecommendationImpactedResourceDismisRequest{
	// ...
}


read, err := client.CreateRecommendationImpactedResourceDismis(ctx, id, payload)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := recommendationimpactedresource.CreateRecommendationImpactedResourcePostponeRequest{
	// ...
}


read, err := client.CreateRecommendationImpactedResourcePostpone(ctx, id, payload)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.CreateRecommendationImpactedResourceReactivate(ctx, id)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.DeleteRecommendationImpactedResource(ctx, id)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.GetRecommendationImpactedResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationImpactedResourceClient.GetRecommendationImpactedResourceCount`

```go
ctx := context.TODO()
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.GetRecommendationImpactedResourceCount(ctx, id)
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
id := recommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

// alternatively `client.ListRecommendationImpactedResources(ctx, id)` can be used to do batched pagination
items, err := client.ListRecommendationImpactedResourcesComplete(ctx, id)
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
id := recommendationimpactedresource.NewDirectoryRecommendationIdImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := recommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.UpdateRecommendationImpactedResource(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
