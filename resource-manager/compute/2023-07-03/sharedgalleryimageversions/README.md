
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2023-07-03/sharedgalleryimageversions` Documentation

The `sharedgalleryimageversions` SDK allows for interaction with the Azure Resource Manager Service `compute` (API Version `2023-07-03`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2023-07-03/sharedgalleryimageversions"
```


### Client Initialization

```go
client := sharedgalleryimageversions.NewSharedGalleryImageVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SharedGalleryImageVersionsClient.Get`

```go
ctx := context.TODO()
id := sharedgalleryimageversions.NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SharedGalleryImageVersionsClient.List`

```go
ctx := context.TODO()
id := sharedgalleryimageversions.NewImageID("12345678-1234-9876-4563-123456789012", "locationValue", "sharedGalleryValue", "imageValue")

// alternatively `client.List(ctx, id, sharedgalleryimageversions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, sharedgalleryimageversions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
