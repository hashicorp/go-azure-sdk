
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation` Documentation

The `recommendation` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation"
```


### Client Initialization

```go
client := recommendation.NewRecommendationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationClient.CreateRecommendation`

```go
ctx := context.TODO()

payload := recommendation.Recommendation{
	// ...
}


read, err := client.CreateRecommendation(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.CreateRecommendationComplete`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.CreateRecommendationComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.CreateRecommendationDismis`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := recommendation.CreateRecommendationDismisRequest{
	// ...
}


read, err := client.CreateRecommendationDismis(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.CreateRecommendationPostpone`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := recommendation.CreateRecommendationPostponeRequest{
	// ...
}


read, err := client.CreateRecommendationPostpone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.CreateRecommendationReactivate`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.CreateRecommendationReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.DeleteRecommendation`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.DeleteRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.GetRecommendation`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.GetRecommendation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.GetRecommendationCount`

```go
ctx := context.TODO()


read, err := client.GetRecommendationCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.ListRecommendations`

```go
ctx := context.TODO()


// alternatively `client.ListRecommendations(ctx)` can be used to do batched pagination
items, err := client.ListRecommendationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RecommendationClient.UpdateRecommendation`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := recommendation.Recommendation{
	// ...
}


read, err := client.UpdateRecommendation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
