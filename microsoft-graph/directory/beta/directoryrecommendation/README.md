
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryrecommendation` Documentation

The `directoryrecommendation` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryrecommendation"
```


### Client Initialization

```go
client := directoryrecommendation.NewDirectoryRecommendationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRecommendationClient.CreateDirectoryRecommendation`

```go
ctx := context.TODO()

payload := directoryrecommendation.Recommendation{
	// ...
}


read, err := client.CreateDirectoryRecommendation(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.CreateDirectoryRecommendationByIdComplete`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.CreateDirectoryRecommendationByIdComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.CreateDirectoryRecommendationByIdDismis`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := directoryrecommendation.CreateDirectoryRecommendationByIdDismisRequest{
	// ...
}


read, err := client.CreateDirectoryRecommendationByIdDismis(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.CreateDirectoryRecommendationByIdPostpone`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := directoryrecommendation.CreateDirectoryRecommendationByIdPostponeRequest{
	// ...
}


read, err := client.CreateDirectoryRecommendationByIdPostpone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.CreateDirectoryRecommendationByIdReactivate`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.CreateDirectoryRecommendationByIdReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.DeleteDirectoryRecommendationById`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.DeleteDirectoryRecommendationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.GetDirectoryRecommendationById`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

read, err := client.GetDirectoryRecommendationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.GetDirectoryRecommendationCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryRecommendationCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRecommendationClient.ListDirectoryRecommendations`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryRecommendations(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryRecommendationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRecommendationClient.UpdateDirectoryRecommendationById`

```go
ctx := context.TODO()
id := directoryrecommendation.NewDirectoryRecommendationID("recommendationIdValue")

payload := directoryrecommendation.Recommendation{
	// ...
}


read, err := client.UpdateDirectoryRecommendationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
