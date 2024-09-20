
## `github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/metadata` Documentation

The `metadata` SDK allows for interaction with Azure Resource Manager `advisor` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/metadata"
```


### Client Initialization

```go
client := metadata.NewMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MetadataClient.RecommendationMetadataGet`

```go
ctx := context.TODO()
id := metadata.NewMetadataID("name")

read, err := client.RecommendationMetadataGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MetadataClient.RecommendationMetadataList`

```go
ctx := context.TODO()


// alternatively `client.RecommendationMetadataList(ctx)` can be used to do batched pagination
items, err := client.RecommendationMetadataListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
