
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleries` Documentation

The `sharedgalleries` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleries"
```


### Client Initialization

```go
client := sharedgalleries.NewSharedGalleriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SharedGalleriesClient.Get`

```go
ctx := context.TODO()
id := sharedgalleries.NewSharedGalleryID("12345678-1234-9876-4563-123456789012", "locationName", "sharedGalleryName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SharedGalleriesClient.List`

```go
ctx := context.TODO()
id := sharedgalleries.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id, sharedgalleries.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, sharedgalleries.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
