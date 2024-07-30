
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/impactedresource` Documentation

The `impactedresource` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/impactedresource"
```


### Client Initialization

```go
client := impactedresource.NewImpactedResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImpactedResourceClient.CreateImpactedResource`

```go
ctx := context.TODO()

payload := impactedresource.ImpactedResource{
	// ...
}


read, err := client.CreateImpactedResource(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.CreateImpactedResourceComplete`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.CreateImpactedResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.CreateImpactedResourceDismis`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := impactedresource.CreateImpactedResourceDismisRequest{
	// ...
}


read, err := client.CreateImpactedResourceDismis(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.CreateImpactedResourcePostpone`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := impactedresource.CreateImpactedResourcePostponeRequest{
	// ...
}


read, err := client.CreateImpactedResourcePostpone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.CreateImpactedResourceReactivate`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.CreateImpactedResourceReactivate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.DeleteImpactedResource`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.DeleteImpactedResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.GetImpactedResource`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

read, err := client.GetImpactedResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.GetImpactedResourceCount`

```go
ctx := context.TODO()


read, err := client.GetImpactedResourceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImpactedResourceClient.ListImpactedResources`

```go
ctx := context.TODO()


// alternatively `client.ListImpactedResources(ctx)` can be used to do batched pagination
items, err := client.ListImpactedResourcesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImpactedResourceClient.UpdateImpactedResource`

```go
ctx := context.TODO()
id := impactedresource.NewDirectoryImpactedResourceID("impactedResourceIdValue")

payload := impactedresource.ImpactedResource{
	// ...
}


read, err := client.UpdateImpactedResource(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
