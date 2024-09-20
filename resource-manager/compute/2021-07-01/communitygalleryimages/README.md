
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimages` Documentation

The `communitygalleryimages` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimages"
```


### Client Initialization

```go
client := communitygalleryimages.NewCommunityGalleryImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CommunityGalleryImagesClient.Get`

```go
ctx := context.TODO()
id := communitygalleryimages.NewCommunityGalleryImageID("12345678-1234-9876-4563-123456789012", "location", "publicGalleryName", "galleryImageName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
