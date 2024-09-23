
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation` Documentation

The `recommendation` SDK allows for interaction with Microsoft Graph `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/recommendation"
```


### Client Initialization

```go
client := recommendation.NewRecommendationClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecommendationClient.CreateRecommendation`

```go
ctx := context.TODO()

payload := recommendation.Recommendation{
	// ...
}


read, err := client.CreateRecommendation(ctx, payload, recommendation.DefaultCreateRecommendationOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

read, err := client.CreateRecommendationComplete(ctx, id, recommendation.DefaultCreateRecommendationCompleteOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

payload := recommendation.CreateRecommendationPostponeRequest{
	// ...
}


read, err := client.CreateRecommendationPostpone(ctx, id, payload, recommendation.DefaultCreateRecommendationPostponeOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

read, err := client.CreateRecommendationReactivate(ctx, id, recommendation.DefaultCreateRecommendationReactivateOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

read, err := client.DeleteRecommendation(ctx, id, recommendation.DefaultDeleteRecommendationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.DismissRecommendation`

```go
ctx := context.TODO()
id := recommendation.NewDirectoryRecommendationID("recommendationId")

payload := recommendation.DismissRecommendationRequest{
	// ...
}


read, err := client.DismissRecommendation(ctx, id, payload, recommendation.DefaultDismissRecommendationOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

read, err := client.GetRecommendation(ctx, id, recommendation.DefaultGetRecommendationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecommendationClient.GetRecommendationsCount`

```go
ctx := context.TODO()


read, err := client.GetRecommendationsCount(ctx, recommendation.DefaultGetRecommendationsCountOperationOptions())
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


// alternatively `client.ListRecommendations(ctx, recommendation.DefaultListRecommendationsOperationOptions())` can be used to do batched pagination
items, err := client.ListRecommendationsComplete(ctx, recommendation.DefaultListRecommendationsOperationOptions())
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
id := recommendation.NewDirectoryRecommendationID("recommendationId")

payload := recommendation.Recommendation{
	// ...
}


read, err := client.UpdateRecommendation(ctx, id, payload, recommendation.DefaultUpdateRecommendationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
