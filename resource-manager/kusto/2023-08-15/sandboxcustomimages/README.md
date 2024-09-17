
## `github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-08-15/sandboxcustomimages` Documentation

The `sandboxcustomimages` SDK allows for interaction with Azure Resource Manager `kusto` (API Version `2023-08-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2023-08-15/sandboxcustomimages"
```


### Client Initialization

```go
client := sandboxcustomimages.NewSandboxCustomImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SandboxCustomImagesClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewKustoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := sandboxcustomimages.SandboxCustomImagesCheckNameRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SandboxCustomImagesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sandboxcustomimages.NewSandboxCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "sandboxCustomImageValue")

payload := sandboxcustomimages.SandboxCustomImage{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SandboxCustomImagesClient.Delete`

```go
ctx := context.TODO()
id := sandboxcustomimages.NewSandboxCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "sandboxCustomImageValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SandboxCustomImagesClient.Get`

```go
ctx := context.TODO()
id := sandboxcustomimages.NewSandboxCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "sandboxCustomImageValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SandboxCustomImagesClient.ListByCluster`

```go
ctx := context.TODO()
id := commonids.NewKustoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

// alternatively `client.ListByCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SandboxCustomImagesClient.Update`

```go
ctx := context.TODO()
id := sandboxcustomimages.NewSandboxCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue", "sandboxCustomImageValue")

payload := sandboxcustomimages.SandboxCustomImage{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
