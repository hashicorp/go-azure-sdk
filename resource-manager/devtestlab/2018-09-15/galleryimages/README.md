
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/galleryimages` Documentation

The `galleryimages` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/galleryimages"
```


### Client Initialization

```go
client := galleryimages.NewGalleryImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GalleryImagesClient.List`

```go
ctx := context.TODO()
id := galleryimages.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

// alternatively `client.List(ctx, id, galleryimages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, galleryimages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
