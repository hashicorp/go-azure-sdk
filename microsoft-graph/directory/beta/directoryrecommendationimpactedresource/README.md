
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryrecommendationimpactedresource` Documentation

The `directoryrecommendationimpactedresource` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryrecommendationimpactedresource"
```


### Client Initialization

```go
client := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.CreateDirectoryRecommendationByIdImpactedResource`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

payload := directoryrecommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.CreateDirectoryRecommendationByIdImpactedResource(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.CreateDirectoryRecommendationByIdImpactedResourceByIdComplete`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.CreateDirectoryRecommendationByIdImpactedResourceByIdComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.CreateDirectoryRecommendationByIdImpactedResourceByIdDismis`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := directoryrecommendationimpactedresource.CreateDirectoryRecommendationByIdImpactedResourceByIdDismisRequest{
	// ...
}


read, err := client.CreateDirectoryRecommendationByIdImpactedResourceByIdDismis(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.CreateDirectoryRecommendationByIdImpactedResourceByIdPostpone`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := directoryrecommendationimpactedresource.CreateDirectoryRecommendationByIdImpactedResourceByIdPostponeRequest{
	// ...
}


read, err := client.CreateDirectoryRecommendationByIdImpactedResourceByIdPostpone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.CreateDirectoryRecommendationByIdImpactedResourceByIdReactivate`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.CreateDirectoryRecommendationByIdImpactedResourceByIdReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.DeleteDirectoryRecommendationByIdImpactedResourceById`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.DeleteDirectoryRecommendationByIdImpactedResourceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.GetDirectoryRecommendationByIdImpactedResourceById`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

read, err := client.GetDirectoryRecommendationByIdImpactedResourceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.GetDirectoryRecommendationByIdImpactedResourceCount`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.GetDirectoryRecommendationByIdImpactedResourceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.ListDirectoryRecommendationByIdImpactedResources`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationID("recommendationIdValue")

// alternatively `client.ListDirectoryRecommendationByIdImpactedResources(ctx, id)` can be used to do batched pagination
items, err := client.ListDirectoryRecommendationByIdImpactedResourcesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRecommendationImpactedResourceClient.UpdateDirectoryRecommendationByIdImpactedResourceById`

```go
ctx := context.TODO()
id := directoryrecommendationimpactedresource.NewDirectoryRecommendationImpactedResourceID("recommendationIdValue", "impactedResourceIdValue")

payload := directoryrecommendationimpactedresource.ImpactedResource{
	// ...
}


read, err := client.UpdateDirectoryRecommendationByIdImpactedResourceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
