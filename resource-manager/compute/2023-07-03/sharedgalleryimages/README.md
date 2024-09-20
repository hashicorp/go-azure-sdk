
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2023-07-03/sharedgalleryimages` Documentation

The `sharedgalleryimages` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2023-07-03`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2023-07-03/sharedgalleryimages"
```


### Client Initialization

```go
client := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SharedGalleryImagesClient.Get`

```go
ctx := context.TODO()
id := sharedgalleryimages.NewImageID("12345678-1234-9876-4563-123456789012", "location", "galleryUniqueName", "galleryImageName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SharedGalleryImagesClient.List`

```go
ctx := context.TODO()
id := sharedgalleryimages.NewSharedGalleryID("12345678-1234-9876-4563-123456789012", "location", "galleryUniqueName")

// alternatively `client.List(ctx, id, sharedgalleryimages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, sharedgalleryimages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
