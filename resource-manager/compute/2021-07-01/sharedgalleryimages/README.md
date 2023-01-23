
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleryimages` Documentation

The `sharedgalleryimages` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleryimages"
```


### Client Initialization

```go
client := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SharedGalleryImagesClient.Get`

```go
ctx := context.TODO()
id := sharedgalleryimages.NewImageID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue")

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
id := sharedgalleryimages.NewSharedGalleryID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue")

// alternatively `client.List(ctx, id, sharedgalleryimages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, sharedgalleryimages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
