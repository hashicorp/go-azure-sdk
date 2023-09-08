
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryimpactedresource` Documentation

The `directoryimpactedresource` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryimpactedresource"
```


### Client Initialization

```go
client := directoryimpactedresource.NewDirectoryImpactedResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryImpactedResourceClient.CreateDirectoryImpactedResource`

```go
ctx := context.TODO()

payload := directoryimpactedresource.ImpactedResource{
	// ...
}


read, err := client.CreateDirectoryImpactedResource(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.CreateDirectoryImpactedResourceByIdComplete`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.CreateDirectoryImpactedResourceByIdComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.CreateDirectoryImpactedResourceByIdDismis`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := directoryimpactedresource.CreateDirectoryImpactedResourceByIdDismisRequest{
	// ...
}


read, err := client.CreateDirectoryImpactedResourceByIdDismis(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.CreateDirectoryImpactedResourceByIdPostpone`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := directoryimpactedresource.CreateDirectoryImpactedResourceByIdPostponeRequest{
	// ...
}


read, err := client.CreateDirectoryImpactedResourceByIdPostpone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.CreateDirectoryImpactedResourceByIdReactivate`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.CreateDirectoryImpactedResourceByIdReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.DeleteDirectoryImpactedResourceById`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.DeleteDirectoryImpactedResourceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.GetDirectoryImpactedResourceById`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.GetDirectoryImpactedResourceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.GetDirectoryImpactedResourceCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryImpactedResourceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryImpactedResourceClient.ListDirectoryImpactedResources`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryImpactedResources(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryImpactedResourcesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryImpactedResourceClient.UpdateDirectoryImpactedResourceById`

```go
ctx := context.TODO()
id := directoryimpactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := directoryimpactedresource.ImpactedResource{
	// ...
}


read, err := client.UpdateDirectoryImpactedResourceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
