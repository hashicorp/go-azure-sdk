
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/extensions` Documentation

The `extensions` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/extensions"
```


### Client Initialization

```go
client := extensions.NewExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExtensionsClient.ExtensionMetadataGet`

```go
ctx := context.TODO()
id := extensions.NewExtensionTypeVersionID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "extensionTypeName", "versionName")

read, err := client.ExtensionMetadataGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExtensionsClient.ExtensionMetadataList`

```go
ctx := context.TODO()
id := extensions.NewPublisherExtensionTypeID("12345678-1234-9876-4563-123456789012", "locationName", "publisherName", "extensionTypeName")

read, err := client.ExtensionMetadataList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExtensionsClient.ExtensionMetadataV2Get`

```go
ctx := context.TODO()
id := extensions.NewVersionID("locationName", "publisherName", "extensionTypeName", "versionName")

read, err := client.ExtensionMetadataV2Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExtensionsClient.ExtensionMetadataV2List`

```go
ctx := context.TODO()
id := extensions.NewExtensionTypeID("locationName", "publisherName", "extensionTypeName")

// alternatively `client.ExtensionMetadataV2List(ctx, id)` can be used to do batched pagination
items, err := client.ExtensionMetadataV2ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExtensionsClient.ExtensionPublisherList`

```go
ctx := context.TODO()
id := extensions.NewLocationID("locationName")

// alternatively `client.ExtensionPublisherList(ctx, id)` can be used to do batched pagination
items, err := client.ExtensionPublisherListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExtensionsClient.ExtensionTypeList`

```go
ctx := context.TODO()
id := extensions.NewPublisherID("locationName", "publisherName")

// alternatively `client.ExtensionTypeList(ctx, id)` can be used to do batched pagination
items, err := client.ExtensionTypeListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
